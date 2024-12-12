// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { useCallback } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { DebugState, TabItem, View, ViewForm } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { base, types } from "../../../wailsjs/go/models";
import { AbisFormDef, AbisTableDef } from "../Abis";
import { NamesFormDef, NamesTableDef } from "../Names";
import { UploadsFormDef, UploadsTableDef } from "../Uploads";
// EXISTING_CODE

export const SharingView = () => {
  const { names, fetchNames, abis, fetchAbis, uploads, fetchUploads } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  const fetchSharing = useCallback(
    (currentItem: number, itemsPerPage: number) => {
      fetchNames(currentItem, itemsPerPage);
      fetchAbis(currentItem, itemsPerPage);
      fetchUploads(currentItem, itemsPerPage);
    },
    [fetchNames, fetchAbis, fetchUploads]
  );

  const namesTable = useReactTable({
    data: names?.items || [],
    columns: NamesTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const abisTable = useReactTable({
    data: abis?.items || [],
    columns: AbisTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const uploadsTable = useReactTable({
    data: uploads?.items || [],
    columns: UploadsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const e =
    names?.items?.length > 0
      ? names.items[0]
      : types.Name.createFrom({
          address: "0xf503017d7baf7fbc0fff7492b751025c6a78179b" as unknown as base.Address,
          name: "string;",
          tags: "string;",
          source: "string;",
          symbol: "string;",
          decimals: 12,
          deleted: false,
        });

  const tabItems: ViewForm = {
    names: <TabItem data={names} groups={NamesFormDef(e, namesTable)} />,
    abis: <TabItem data={abis} groups={AbisFormDef(abisTable)} />,
    uploads: <TabItem data={uploads} groups={UploadsFormDef(uploadsTable)} />,
  };

  // if (!(status?.items?.length > 0)) {
  //   return <>Loading...</>;
  // }

  return (
    <ViewStateProvider
      // do not remove - delint
      nItems={names.nItems}
      fetchFn={fetchSharing}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={[names.updater, abis.updater, uploads.updater]} />
      <View tabItems={tabItems} searchable />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
