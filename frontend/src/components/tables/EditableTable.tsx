import { useState } from "react";
import { useReactTable, getCoreRowModel, ColumnDef, Table } from "@tanstack/react-table";

type EditableTableProps<T> = {
  data: T[];
  columns: ColumnDef<T>[];
};

export const EditableTable = <T,>({ data: initialData, columns }: EditableTableProps<T>) => {
  const [data, setData] = useState(initialData);
  const [editingRowId, setEditingRowId] = useState<number | null>(null);
  const [editingData, setEditingData] = useState<Partial<T>>({});

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>, field: keyof T) => {
    setEditingData({
      ...editingData,
      [field]: e.target.value,
    });
  };

  const toggleEdit = (rowId: number) => {
    if (editingRowId === rowId) {
      setData((prev) => prev.map((row, index) => (index === rowId ? { ...row, ...editingData } : row)));
      setEditingRowId(null);
      setEditingData({});
    } else {
      setEditingRowId(rowId);
      const rowData = data[rowId];
      setEditingData(rowData ?? {});
    }
  };

  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <MyTable
      table={table}
      toggleEdit={toggleEdit}
      handleInputChange={handleInputChange}
      editingRowId={editingRowId}
      editingData={editingData}
    />
  );
};

const MyTable = <T,>({
  table,
  toggleEdit,
  handleInputChange,
  editingRowId,
  editingData,
}: {
  table: Table<T>;
  toggleEdit: (id: number) => void;
  handleInputChange: (e: React.ChangeEvent<HTMLInputElement>, field: keyof T) => void;
  editingRowId: number | null;
  editingData: Partial<T>;
}) => {
  return (
    <table style={{ width: "100%", borderCollapse: "collapse", border: "1px solid black" }}>
      <MyTableHeader table={table} />
      <MyTableBody
        table={table}
        toggleEdit={toggleEdit}
        handleInputChange={handleInputChange}
        editingRowId={editingRowId}
        editingData={editingData}
      />
    </table>
  );
};

const MyTableHeader = <T,>({ table }: { table: Table<T> }) => {
  return (
    <thead>
      {table.getHeaderGroups().map((headerGroup) => (
        <tr key={headerGroup.id}>
          {headerGroup.headers.map((header) => (
            <th key={header.id}>
              {header.isPlaceholder
                ? null
                : typeof header.column.columnDef.header === "function"
                  ? header.column.columnDef.header(header.getContext())
                  : header.column.columnDef.header}
            </th>
          ))}
        </tr>
      ))}
    </thead>
  );
};

const MyTableBody = <T,>({
  table,
  toggleEdit,
  handleInputChange,
  editingRowId,
  editingData,
}: {
  table: Table<T>;
  toggleEdit: (id: number) => void;
  handleInputChange: (e: React.ChangeEvent<HTMLInputElement>, field: keyof T) => void;
  editingRowId: number | null;
  editingData: Partial<T>;
}) => {
  return (
    <tbody>
      {table.getRowModel().rows.map((row, rowIndex) => (
        <tr key={row.id}>
          {row.getVisibleCells().map((cell) => (
            <td key={cell.id}>
              {editingRowId === rowIndex ? (
                <input
                  value={
                    typeof editingData[cell.column.id as keyof T] === "string" ||
                    typeof editingData[cell.column.id as keyof T] === "number"
                      ? (editingData[cell.column.id as keyof T] as string | number)
                      : ""
                  }
                  onChange={(e) => handleInputChange(e, cell.column.id as keyof T)}
                />
              ) : typeof cell.column.columnDef.cell === "function" ? (
                cell.column.columnDef.cell(cell.getContext())
              ) : (
                cell.getValue()
              )}
            </td>
          ))}
          <td>
            <button style={{ color: "darkblue", border: "1px solid darkblue" }} onClick={() => toggleEdit(rowIndex)}>
              {editingRowId === rowIndex ? "Save" : "Edit"}
            </button>
          </td>
        </tr>
      ))}
    </tbody>
  );
};
