import React from "react";
import "./DataTable.css";
import { Table, Title } from "@mantine/core";
import { flexRender, Table as ReactTable } from "@tanstack/react-table";
import { CustomMeta, Paginator } from "./";
import { useViewState } from "@state";

interface DataTableProps<T> {
  table: ReactTable<T>;
  loading: boolean;
}

export function DataTable<T>({ table, loading }: DataTableProps<T>) {
  const { pager } = useViewState();

  if (loading) {
    return <Title order={3}>Loading...</Title>;
  }

  return (
    <>
      <Table>
        <TableHeader table={table} />
        <TableBody table={table} />
      </Table>
      <Paginator pager={pager} />
    </>
  );
}

interface TablePartProps<T> {
  table: ReactTable<T>;
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

function TableBody<T>({ table }: TablePartProps<T>) {
  return (
    <Table.Tbody>
      {table.getRowModel().rows.map((row) => (
        <Table.Tr key={row.id}>
          {row.getVisibleCells().map((cell) => {
            const meta = cell.column.columnDef.meta as CustomMeta;
            return (
              <Table.Td key={cell.id} className={meta?.className}>
                {flexRender(cell.column.columnDef.cell, cell.getContext())}
              </Table.Td>
            );
          })}
        </Table.Tr>
      ))}
    </Table.Tbody>
  );
}
