// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { useCallback } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { DebugState, TabItem, View, ViewForm } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { AbisFormDef, AbisTableDef } from "../Abis";
import { NamesFormDef, NamesTableDef } from "../Names";
import { PinsFormDef, PinsTableDef } from "../Pins";
import { UploadsFormDef, UploadsTableDef } from "../Uploads";
// EXISTING_CODE

export const SharingView = () => {
  const { names, fetchNames, abis, fetchAbis, pins, fetchPins, uploads, fetchUploads } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  const fetchSharing = useCallback(
    (currentItem: number, itemsPerPage: number) => {
      fetchNames(currentItem, itemsPerPage);
      fetchAbis(currentItem, itemsPerPage);
      fetchPins(currentItem, itemsPerPage);
      fetchUploads(currentItem, itemsPerPage);
    },
    [fetchNames, fetchAbis, fetchPins, fetchUploads]
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

  const pinsTable = useReactTable({
    data: pins?.items || [],
    columns: PinsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const uploadsTable = useReactTable({
    data: uploads?.items || [],
    columns: UploadsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const tabItems: ViewForm = {
    names: <TabItem data={names} groups={NamesFormDef(namesTable)} />,
    abis: <TabItem data={abis} groups={AbisFormDef(abisTable)} />,
    pins: <TabItem data={pins} groups={PinsFormDef(pinsTable)} />,
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
      <DebugState u={[names.updater, abis.updater, pins.updater, uploads.updater]} />
      <View tabItems={tabItems} searchable />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
