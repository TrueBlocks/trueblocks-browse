import { Group } from "@mantine/core";
import { createColumnHelper } from "@tanstack/react-table";
import {
  CustomColumnDef,
  Formatter,
  ExploreButton,
  ExportButton,
  ViewButton,
  DeleteButton,
  DalleButton,
  GoogleButton,
} from "@components";
import { types } from "@gocode/models";

const columnHelper = createColumnHelper<types.HistoryContainer>();

const baseColumns: CustomColumnDef<types.HistoryContainer, any>[] = [
  // columnHelper.accessor("address", {
  //   header: () => "Dalle",
  //   cell: (info) => <Formatter type="dalle-small" value={info.renderValue()} />,
  //   meta: { className: "medium cell" },
  // }),
  columnHelper.accessor("address", {
    header: () => "Address",
    cell: (info) => <Formatter type="address-editor" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("nItems", {
    header: () => "Loaded",
    cell: (info) => {
      const { nItems, nTotal } = info.row.original;
      return <Formatter type="progress" value={nItems} value2={nTotal} />;
    },
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("balance", {
    header: () => "Balance",
    cell: (info) => <Formatter type="ether" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
];

export const withDelete: CustomColumnDef<types.HistoryContainer, any>[] = [
  ...baseColumns,
  columnHelper.accessor("address", {
    header: () => " ",
    cell: (info) => {
      const { address } = info.row.original;
      const addr = address as unknown as string;
      return (
        <Group wrap={"nowrap"}>
          <ExploreButton noText value={info.renderValue()} />
          <DalleButton noText value={info.renderValue()} />
          <GoogleButton noText value={info.renderValue()} />
          <ViewButton noText value={info.renderValue()} />
          <ExportButton noText value={info.renderValue()} />
          <DeleteButton value={addr} isDeleted={false} />
        </Group>
      );
    },
    meta: { className: "wide cell" },
  }),
];

export const withoutDelete: CustomColumnDef<types.HistoryContainer, any>[] = [
  ...baseColumns,
  columnHelper.accessor("address", {
    header: () => " ",
    cell: (info) => {
      return (
        <Group wrap={"nowrap"}>
          <ExploreButton noText value={info.renderValue()} />
          <DalleButton noText value={info.renderValue()} />
          <GoogleButton noText value={info.renderValue()} />
          <ViewButton noText value={info.renderValue()} />
          <ExportButton noText value={info.renderValue()} />
        </Group>
      );
    },
    meta: { className: "wide cell" },
  }),
];
