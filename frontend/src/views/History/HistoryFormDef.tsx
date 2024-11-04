import { useState, useEffect } from "react";
import { Table } from "@tanstack/react-table";
import { ExploreButton, ExportButton, DataTable, FieldGroup, GoogleButton } from "@components";
import { types, base } from "@gocode/models";
import { GetName } from "../../../wailsjs/go/app/App";
import { useUtils } from "../../hooks";

export const HistoryFormDef = (
  address: base.Address,
  table: Table<types.Transaction>
): FieldGroup<types.HistoryContainer>[] => {
  const [named, setNamed] = useState(address as unknown as string);
  const { ShortenAddr } = useUtils();

  useEffect(() => {
    GetName(address).then((name) => {
      setNamed(name === "" ? ShortenAddr(address as unknown as string) : name);
    });
  }, [address, ShortenAddr]);

  return [
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
  ];
};
