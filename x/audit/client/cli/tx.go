package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/mattverse/errata/x/audit/types"
)

// GetTxCmd returns a root CLI command handler for all x/bond transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Audit transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(NewMsgRegisterProtocolCmd())
	txCmd.AddCommand(NewJoinAttackPoolCmd())
	txCmd.AddCommand(NewJoinDefensePoolCmd())
	txCmd.AddCommand(NewAddErrataCmd())

	return txCmd
}

// NewMsgRegisterProtocolCmd returns a CLI command handler for creating a MsgBondIn transaction.
func NewMsgRegisterProtocolCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-protocol [title] [description] [source-code] [project-home] [category]",
		Short: "register a new protocol",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterProtocol(args[0], args[1], args[2], args[3], args[4], clientCtx.GetFromAddress())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewJoinAttackPoolCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "join-attack-pool [poolId] [tokenIn]",
		Short: "join attack pool",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			tokenIn, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("invalid token amount")
			}

			msg := types.NewMsgJoinAttackPool(uint64(poolID), tokenIn)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewJoinDefensePoolCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "join-defense-pool [poolId] [tokenIn]",
		Short: "join defense pool",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			tokenIn, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("invalid token amount")
			}

			msg := types.NewMsgJoinDefensePool(uint64(poolID), tokenIn)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewAddErrataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-errata [poolId] [vulnerability-type] [errata-code] [vulnerability]",
		Short: "join defense pool",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgAddErrata(uint64(poolID), args[1], args[2], args[3])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
