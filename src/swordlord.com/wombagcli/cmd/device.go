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
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Add, change and manage devices of your users.",
	Long: `Add, change and manage devices of your users. Requires a subcommand.`,
	RunE: nil,
}

var deviceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all devices.",
	Long: `List all devices.`,
	RunE: ListDevices,
}

var deviceAddCmd = &cobra.Command{
	Use:   "add [device] [password] [user]",
	Short: "Add new device to given user.",
	Long: `Add new device to given user.`,
	RunE: AddDevice,
}

var deviceUpdateCmd = &cobra.Command{
	Use:   "update [device] [password]",
	Short: "Update the password of the device.",
	Long: `Update the password of the device.`,
	RunE: UpdateDevice,
}

var deviceDeleteCmd = &cobra.Command{
	Use:   "delete [device]",
	Short: "Deletes a device.",
	Long: `Deletes a device.`,
	RunE: DeleteDevice,
}

func ListDevices(cmd *cobra.Command, args []string) error {

	tablemodule.ListDevice()

	return nil
}

func AddDevice(cmd *cobra.Command, args []string) error {

	if len(args) != 3 {
		er("command 'add' needs a device, password and user")
	} else {
		tablemodule.AddDevice(args[0], args[1], args[2])
	}

	return nil
}

func UpdateDevice(cmd *cobra.Command, args []string) error {

	if len(args) != 2 {
		er("command 'update' needs a device and a new password")
	} else {
		tablemodule.UpdateDevice(args[0], args[1])
	}

	return nil
}

func DeleteDevice(cmd *cobra.Command, args []string) error {

	if len(args) < 1 {
		er("command 'delete' needs a device Id")
	} else {
		tablemodule.DeleteDevice(args[0])
	}

	return nil
}

func init() {
	RootCmd.AddCommand(deviceCmd)

	deviceCmd.AddCommand(deviceListCmd)
	deviceCmd.AddCommand(deviceAddCmd)
	deviceCmd.AddCommand(deviceUpdateCmd)
	deviceCmd.AddCommand(deviceDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
