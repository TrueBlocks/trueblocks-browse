import React from "react";
import "./DataTable.css";
import { Table, Title, Button, Box } from "@mantine/core";
import { flexRender, Table as ReactTable } from "@tanstack/react-table";
import { CustomMeta, Paginator } from "./";
import { useViewState } from "@state";

interface DataTableProps<T> {
  table: ReactTable<T>;
  bumper?: boolean;
  loading: boolean;
}

export function DataTable<T>({ table, bumper, loading }: DataTableProps<T>) {
  const { pager } = useViewState();

  if (loading) {
    return <Title order={3}>Loading...</Title>;
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
  return (
    <Table.Tbody>
      {table.getRowModel().rows.map((row, index) => (
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
      ))}
    </Table.Tbody>
  );
}
