// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { createColumnHelper } from "@tanstack/react-table";
import { types } from "@gocode/models";
import { CustomColumnDef, Formatter } from "../../components";
// EXISTING_CODE

const columnHelper = createColumnHelper<types.Chain>();

export const ConfigTableDef: CustomColumnDef<types.Chain, any>[] = [
  // EXISTING_CODE
  columnHelper.accessor("chain", {
    header: () => "Chain",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("chainId", {
    header: () => "ChainId",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("symbol", {
    header: () => "Symbol",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("rpcProvider", {
    header: () => "RpcProvider",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("localExplorer", {
    header: () => "LocalExplorer",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("remoteExplorer", {
    header: () => "RemoteExplorer",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("ipfsGateway", {
    header: () => "IpfsGateway",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("keyEndpoint", {
    header: () => "KeyEndpoint",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  // columnHelper.accessor("scrape", {
  //   header: () => "Scrape",
  //   cell: (info) => <Formatter type="text" value={info.renderValue()} />,
  //   meta: { className: "medium cell" },
  // }),
  // EXISTING_CODE
];

// EXISTING_CODE
// EXISTING_CODE
