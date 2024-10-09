import "./DataTable.css";
import { Table, Title, Box, Alert } from "@mantine/core";
import { IconInfoCircle } from "@tabler/icons-react";
import { flexRender, Table as ReactTable } from "@tanstack/react-table";
import { useViewState } from "@state";
import { CustomMeta, Paginator } from "./";

interface DataTableProps<T> {
  table: ReactTable<T>;
  loading: boolean;
}

export function DataTable<T>({ table, loading }: DataTableProps<T>) {
  const { pager, nItems } = useViewState();

  if (loading) {
    return <Title order={3}>Loading...</Title>;
  }

  if (nItems <= 0) {
    return (
      <Box style={{ width: "100%" }}>
        <Alert variant="light" color="blue" title="No data found" icon={<IconInfoCircle />} style={{ width: "100%" }} />
      </Box>
    );
  }

  const selectedRow = pager.selected % pager.perPage;
  return (
    <>
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
      {table.getHeaderGroups().map((headerGroup, headerGroupIndex) => (
        <Table.Tr key={`headerGroup-${headerGroupIndex}`}>
          {headerGroup.headers.map((header, headerIndex) => (
            <Table.Th key={`header-${header.id}-${headerIndex}`} className="centered">
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
  const inner = table.getRowModel().rows.map((row, rowIndex) => {
    const rowKey = `row-${rowIndex}-${row.id}`;
    return (
      <Table.Tr
        key={rowKey}
        className={rowIndex === selectedRow ? "selected-row" : ""}
        onClick={() => pager.setRecord(pager.getOffset() + rowIndex)}
      >
        {row.getVisibleCells().map((cell, cellIndex) => {
          const cellKey = `cell-${rowIndex}-${cellIndex}-${cell.id}`;
          const meta = cell.column.columnDef.meta as CustomMeta;
          return (
            <Table.Td key={cellKey} className={meta?.className}>
              {flexRender(cell.column.columnDef.cell, cell.getContext())}
            </Table.Td>
          );
        })}
      </Table.Tr>
    );
  });
  return <Table.Tbody>{inner}</Table.Tbody>;
}
