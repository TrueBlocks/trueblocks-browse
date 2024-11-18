// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter } from "@components";
import { types } from "@gocode/models";
// EXISTING_CODE

const columnHelper = createColumnHelper<types.CacheItem>();

export const StatusTableDef: CustomColumnDef<types.CacheItem, any>[] = [
  // EXISTING_CODE
  columnHelper.accessor("type", {
    header: () => "Type",
    cell: (info) => <Formatter type="text" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("nFolders", {
    header: () => "nFolders",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("nFiles", {
    header: () => "nFiles",
    cell: (info) => <Formatter type="int" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("sizeInBytes", {
    header: () => "SizeInBytes",
    cell: (info) => <Formatter type="bytes" value={info.renderValue()} />,
    meta: { className: "small cell" },
  }),
  columnHelper.accessor("lastCached", {
    header: () => "LastCached",
    cell: (info) => <Formatter type="date" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("path", {
    header: () => "Path",
    cell: (info) => <Formatter type="path" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
  // EXISTING_CODE
];

// EXISTING_CODE
// EXISTING_CODE

//-------------------------------------------------------------------
// Template variables:
// class:         Status
// lower:         status
// routeLabel:    Status
// routeLower:    status
// itemName:      CacheItem
// embedType:     coreTypes.Status
