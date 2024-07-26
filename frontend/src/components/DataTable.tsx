import { useRef } from "react";
import "./DataTable.css";
import { Card, Table, Title } from "@mantine/core";
import { flexRender, Table as ReactTable } from "@tanstack/react-table";
import { usePopup } from "./DataTablePopup";
import { CustomMeta } from "./";

export function DataTable<T>({ table, loading }: { table: ReactTable<T>; loading: boolean }) {
  const popup = useRef<HTMLDivElement>(null);
  const { floatingStyles, setTarget, triggerProps } = usePopup(popup);

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
                    <Table.Td key={cell.id} className={meta?.className} onClick={(e) => setTarget(e.currentTarget)} {...triggerProps}>
                        {flexRender(cell.column.columnDef.cell, cell.getContext())}
                    </Table.Td>
                  );
                })}
              </Table.Tr>
            ))}
          </Table.Tbody>
        </Table>


        <div ref={popup} style={floatingStyles}>
          <Card shadow="sm" padding="sm" withBorder>
            Popover!
          </Card>
        </div>

      </>
    );
  }
}
