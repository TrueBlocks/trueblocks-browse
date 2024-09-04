import React, { useEffect, useMemo } from "react";

import "./DataTable.css";
import { Table, Title, Box, Alert } from "@mantine/core";
import { useViewState } from "@state";
import { IconInfoCircle } from "@tabler/icons-react";
import { flexRender, Table as ReactTable } from "@tanstack/react-table";

import { CustomMeta, Paginator } from "./";

interface DataTableProps<T> {
  table: ReactTable<T>;
  bumper?: boolean;
  loading: boolean;
}

export function DataTable<T>({ table, bumper, loading }: DataTableProps<T>) {
  const { pager, nItems } = useViewState();

  if (loading) {
    return <Title order={3}>Loading...</Title>;
  }

  if (nItems <= 0) {
    return <Alert variant="light" color="blue" title="No data found" icon={<IconInfoCircle />}></Alert>;
  }

  const selectedRow = pager.selected % pager.perPage;
  return (
    <>
      {bumper ? <Box className="bumper">{JSON.stringify(pager, null, 2)}</Box> : <></>}
      <Table>
        <TableHeader table={table} />
        <TableBody table={table} selectedRow={selectedRow} />
      </Table>
      <Paginator pager={pager} />
    </>
  );
}

interface TablePartProps<T> {
  table: ReactTable<T>;
  selectedRow?: number;
}

function TableHeader<T>({ table }: TablePartProps<T>) {
  return (
    <Table.Thead>
      {table.getHeaderGroups().map((headerGroup) => (
        <Table.Tr key={headerGroup.id}>
          {headerGroup.headers.map((header) => (
            <Table.Th key={header.id} className="centered">
              {header.isPlaceholder ? null : flexRender(header.column.columnDef.header, header.getContext())}
            </Table.Th>
          ))}
        </Table.Tr>
      ))}
    </Table.Thead>
  );
}

function TableBody<T>({ table, selectedRow }: TablePartProps<T>) {
  const { pager } = useViewState(); // Access pager to use setSelected
  const inner = useMemo(() => {
    return table.getRowModel().rows.map((row, index) => (
      <Table.Tr
        key={row.id}
        className={index === selectedRow ? "selected-row" : ""}
        onClick={() => pager.setRecord(index + pager.getOffset())}
      >
        {row.getVisibleCells().map((cell) => {
          const meta = cell.column.columnDef.meta as CustomMeta;
          return (
            <Table.Td key={cell.id} className={meta?.className}>
              {flexRender(cell.column.columnDef.cell, cell.getContext())}
            </Table.Td>
          );
        })}
      </Table.Tr>
    ));
  }, [pager, selectedRow, table]);
  return <Table.Tbody>{inner}</Table.Tbody>;
}
