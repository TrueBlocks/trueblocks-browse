import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";
import { types } from "@gocode/models";

const columnHelper = createColumnHelper<types.ChunkStats>();

export const IndexesTableDef: CustomColumnDef<types.ChunkStats, any>[] = [
  columnHelper.accessor("range", {
    header: () => "range",
    cell: (info) => <Formatter type="range" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nBlocks", {
    header: () => "nBlocks",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("nAddrs", {
    header: () => "nAddrs",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("nApps", {
    header: () => "nApps",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("chunkSz", {
    header: () => "chunkSz",
    cell: (info) => <Formatter type="bytes" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("nBlooms", {
    header: () => "nBlooms",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("bloomSz", {
    header: () => "bloomSz",
    cell: (info) => <Formatter type="bytes" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("addrsPerBlock", {
    header: () => "addrsPerBlock",
    cell: (info) => <Formatter type="float" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("appsPerAddr", {
    header: () => "appsPerAddr",
    cell: (info) => <Formatter type="float" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("appsPerBlock", {
    header: () => "appsPerBlock",
    cell: (info) => <Formatter type="float" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
];
