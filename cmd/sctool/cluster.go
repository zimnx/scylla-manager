// Copyright (C) 2017 ScyllaDB

package main

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/scylladb/mermaid/mermaidclient"
	"github.com/spf13/cobra"
)

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Add or delete clusters",
}

func init() {
	register(clusterCmd, rootCmd)
}

var (
	cfgClusterName            string
	cfgClusterHost            string
	cfgClusterSSHUser         string
	cfgClusterSSHIdentityFile string
)

func clusterInitCommonFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&cfgClusterName, "name", "n", "", "`alias` you can give to your cluster")
	cmd.Flags().StringVar(&cfgClusterHost, "host", "", "hostname or IP of one of the cluster nodes")
	cmd.Flags().StringVar(&cfgClusterSSHUser, "ssh-user", "", "SSH user `name` used to connect to the cluster nodes")
	cmd.Flags().StringVar(&cfgClusterSSHIdentityFile, "ssh-identity-file", "", "`path` to identity file containing SSH private key")
}

var clusterAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a cluster to manager",

	RunE: func(cmd *cobra.Command, args []string) error {
		c := &mermaidclient.Cluster{
			Name: cfgClusterName,
			Host: cfgClusterHost,
		}

		if cfgClusterSSHUser != "" && cfgClusterSSHIdentityFile == "" {
			return printableError{errors.New("missing flag \"ssh-identity-file\"")}
		}
		if cfgClusterSSHIdentityFile != "" && cfgClusterSSHUser == "" {
			return printableError{errors.New("missing flag \"ssh-user\"")}
		}
		if cfgClusterSSHUser != "" && cfgClusterSSHIdentityFile != "" {
			b, err := ioutil.ReadFile(cfgClusterSSHIdentityFile)
			if err != nil {
				return printableError{inner: err}
			}
			c.SSHIdentityFile = b
			c.SSHUser = cfgClusterSSHUser
		}

		id, err := client.CreateCluster(ctx, c)
		if err != nil {
			return printableError{err}
		}

		w := cmd.OutOrStdout()
		fmt.Fprintln(w, id)

		tasks, err := client.ListTasks(ctx, id, "repair", false, "")
		if err != nil {
			return printableError{err}
		}
		if len(tasks) > 0 {
			s := tasks[0].Schedule
			w := cmd.OutOrStderr()
			fmt.Fprintf(w, clipper, id, formatTime(s.StartDate), s.Interval, id)
		}

		return nil
	},
}

func init() {
	cmd := clusterAddCmd
	register(cmd, clusterCmd)

	clusterInitCommonFlags(cmd)
	requireFlags(cmd, "host")
}

var clusterUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Modifies a cluster",

	RunE: func(cmd *cobra.Command, args []string) error {
		cluster, err := client.GetCluster(ctx, cfgCluster)
		if err != nil {
			return printableError{err}
		}

		ok := false
		if cmd.Flags().Changed("name") {
			cluster.Name = cfgClusterName
			ok = true
		}
		if cmd.Flags().Changed("host") {
			cluster.Host = cfgClusterHost
			ok = true
		}
		if cmd.Flags().Changed("ssh-user") {
			cluster.SSHUser = cfgClusterSSHUser
			ok = true
		}
		if cmd.Flags().Changed("ssh-identity-file") {
			b, err := ioutil.ReadFile(cfgClusterSSHIdentityFile)
			if err != nil {
				return printableError{inner: err}
			}
			cluster.SSHIdentityFile = b
			ok = true
		}
		if !ok {
			return errors.New("nothing to do")
		}

		if err := client.UpdateCluster(ctx, cluster); err != nil {
			return printableError{err}
		}

		return nil
	},
}

func init() {
	cmd := clusterUpdateCmd
	register(cmd, clusterCmd)

	clusterInitCommonFlags(cmd)
}

var clusterDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a cluster from manager",

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := client.DeleteCluster(ctx, cfgCluster); err != nil {
			return printableError{err}
		}

		return nil
	},
}

func init() {
	cmd := clusterDeleteCmd
	register(cmd, clusterCmd)
}

var clusterListCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows managed clusters",

	RunE: func(cmd *cobra.Command, args []string) error {
		clusters, err := client.ListClusters(ctx)
		if err != nil {
			return printableError{err}
		}

		t := newTable("cluster id", "name", "host", "ssh user")
		for _, c := range clusters {
			t.AddRow(c.ID, c.Name, c.Host, c.SSHUser)
		}
		fmt.Fprint(cmd.OutOrStdout(), t)

		return nil
	},
}

func init() {
	register(clusterListCmd, clusterCmd)
}
