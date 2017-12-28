package cmd

/*-----------------------------------------------------------------------------
 **
 ** - Wombag -
 **
 ** the alternative, native backend for your Wallabag apps
 **
 ** Copyright 2017 by SwordLord - the coding crew - http://www.swordlord.com
 ** and contributing authors
 **
 ** This program is free software; you can redistribute it and/or modify it
 ** under the terms of the GNU Affero General Public License as published by the
 ** Free Software Foundation, either version 3 of the License, or (at your option)
 ** any later version.
 **
 ** This program is distributed in the hope that it will be useful, but WITHOUT
 ** ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 ** FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License
 ** for more details.
 **
 ** You should have received a copy of the GNU Affero General Public License
 ** along with this program. If not, see <http://www.gnu.org/licenses/>.
 **
 **-----------------------------------------------------------------------------
 **
 ** Original Authors:
 ** LordEidi@swordlord.com
 ** LordLightningBolt@swordlord.com
 **
-----------------------------------------------------------------------------*/
import (
	"github.com/spf13/cobra"
	"swordlord.com/wombag/tablemodule"
)

// domainCmd represents the domain command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Add, change and manage your users.",
	Long: `Add, change and manage your users. Requires a subcommand.`,
	RunE: nil,
}

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users.",
	Long: `List all users.`,
	RunE: ListUser,
}

var userAddCmd = &cobra.Command{
	Use:   "add [user] [password]",
	Short: "Add new user to this instance of gohjasmin.",
	Long: `Add new user to this instance of gohjasmin.`,
	RunE: AddUser,
}

var userUpdateCmd = &cobra.Command{
	Use:   "update [user] [password]",
	Short: "Update the password of the user.",
	Long: `Update the password of the user.`,
	RunE: UpdateUser,
}

var userDeleteCmd = &cobra.Command{
	Use:   "delete [user]",
	Short: "Deletes a user and all of her devices.",
	Long: `Deletes a user and all of his or her devices.`,
	RunE: DeleteUser,
}

func ListUser(cmd *cobra.Command, args []string) error {

	tablemodule.ListUser()

	return nil
}

func AddUser(cmd *cobra.Command, args []string) error {

	if len(args) < 2 {
		er("command 'add' needs a user and a password")
	} else {
		tablemodule.AddUser(args[0], args[1])
	}

	return nil
}

func UpdateUser(cmd *cobra.Command, args []string) error {

	if len(args) < 2 {
		er("command 'update' needs a user and a new password")
	} else {
		tablemodule.UpdateUser(args[0], args[1])
	}

	return nil
}

func DeleteUser(cmd *cobra.Command, args []string) error {

	if len(args) < 1 {
		er("command 'delete' needs a user")
	} else {
		tablemodule.DeleteUser(args[0])
	}

	return nil
}

func init() {
	RootCmd.AddCommand(userCmd)

	userCmd.AddCommand(userListCmd)
	userCmd.AddCommand(userAddCmd)
	userCmd.AddCommand(userUpdateCmd)
	userCmd.AddCommand(userDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
