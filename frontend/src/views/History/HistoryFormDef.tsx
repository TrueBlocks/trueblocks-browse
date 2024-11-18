// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
// EXISTING_CODE
import { useState, useEffect } from "react";
import { Table } from "@tanstack/react-table";
import { ExploreButton, ExportButton, DataTable, FieldGroup, GoogleButton } from "@components";
import { GetName } from "@gocode/app/App";
import { types, base } from "@gocode/models";
import { useUtils } from "@hooks";
// EXISTING_CODE

export const HistoryFormDef = (
  table: Table<types.Transaction>,
  address: base.Address
): FieldGroup<types.HistoryContainer>[] => {
  const [named, setNamed] = useState(address as unknown as string);
  const { ShortenAddr } = useUtils();
  useEffect(() => {
    GetName(address).then((name) => {
      setNamed(name === "" ? ShortenAddr(address as unknown as string) : name);
    });
  }, [address, ShortenAddr]);
  return [
    // EXISTING_CODE
    {
      label: "DalleDress",
      colSpan: 2,
      fields: [{ label: "", type: "dalle", accessor: "address" }],
    },
    {
      label: "Transaction Data",
      colSpan: 7,
      fields: [
        { label: "name", type: "address-name-only", accessor: "address" },
        { label: "address", type: "address-address-only", accessor: "address" },
        { label: "balance", type: "ether", accessor: "balance" },
      ],
    },
    {
      label: "Transaction Data",
      colSpan: 3,
      fields: [
        { label: "nTransactions", type: "int", accessor: "nItems" },
        { label: "nLogs", type: "int", accessor: "nLogs" },
        { label: "nTokens", type: "int", accessor: "nTokens" },
        { label: "nErrors", type: "int", accessor: "nErrors" },
      ],
    },
    {
      label: "Buttons",
      buttons: [
        <ExploreButton key={"explore"} value={address} />,
        <GoogleButton key={"google"} value={address} />,
        <ExportButton key={"export"} value={address} />,
      ],
    },
    {
      label: named,
      collapsable: false,
      components: [<DataTable<types.Transaction> key={"dataTable"} table={table} loading={false} />],
    },
    // EXISTING_CODE
  ];
};

// EXISTING_CODE
// EXISTING_CODE

//-------------------------------------------------------------------
// Template variables:
// class:         History
// routeLabel:    History
// itemName:      Transaction
// isHistory:     true
// isSession:     false
// isConfig:      false
