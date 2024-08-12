import React from "react";
import "./DataTable.css";
import { Table, Title } from "@mantine/core";
import { flexRender, Table as ReactTable } from "@tanstack/react-table";
import { DataPopover } from "../popovers";
import { CustomMeta, Paginator, Pager } from "./";

interface DataTableProps<T> {
  table: ReactTable<T>;
  loading: boolean;
  pager: Pager;
}

export function DataTable<T>({ table, loading, pager }: DataTableProps<T>) {
  if (loading) {
    return <Title order={3}>Loading...</Title>;
  } else {
    return (
      <>
        <Table>
          <Table.Thead>
            {table.getHeaderGroups().map((headerGroup) => (
              <Table.Tr key={headerGroup.id}>
                {headerGroup.headers.map((header) => (
                  <Table.Th key={header.id} className={"centered"}>
                    {header.isPlaceholder ? null : flexRender(header.column.columnDef.header, header.getContext())}
                  </Table.Th>
                ))}
              </Table.Tr>
            ))}
          </Table.Thead>
          <Table.Tbody>
            {table.getRowModel().rows.map((row) => (
              <Table.Tr key={row.id}>
                {row.getVisibleCells().map((cell) => {
                  const meta = cell.column.columnDef.meta as CustomMeta;
                  return (
                    <Table.Td key={cell.id} className={meta?.className}>
                      <DataPopover editor={meta.editor?.(cell.getValue)}>
                        {flexRender(cell.column.columnDef.cell, cell.getContext())}
                      </DataPopover>
                    </Table.Td>
                  );
                })}
              </Table.Tr>
            ))}
          </Table.Tbody>
        </Table>
        <Paginator pager={pager} />
      </>
    );
  }
}
