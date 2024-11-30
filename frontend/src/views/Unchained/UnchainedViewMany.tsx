// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { useCallback } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { DebugState, FormTable, View, ViewForm } from "../../components";
import { IndexesFormDef, IndexesTableDef } from "../Indexes";
import { ManifestsFormDef, ManifestsTableDef } from "../Manifests";

// EXISTING_CODE

export const UnchainedView = () => {
  const { indexes, fetchIndexes, manifests, fetchManifests } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  // EXISTING_CODE
  // EXISTING_CODE

  const fetchUnchained = useCallback(
    (currentItem: number, itemsPerPage: number) => {
      fetchIndexes(currentItem, itemsPerPage);
      fetchManifests(currentItem, itemsPerPage);
    },
    [fetchIndexes, fetchManifests]
  );

  const indexesTable = useReactTable({
    data: indexes?.items || [],
    columns: IndexesTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const manifestsTable = useReactTable({
    data: manifests?.items || [],
    columns: ManifestsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "unchained";
  const tabs = ["indexes", "manifests"];
  const forms: ViewForm = {
    indexes: <FormTable data={indexes} groups={IndexesFormDef(indexesTable)} />,
    manifests: <FormTable data={manifests} groups={ManifestsFormDef(manifestsTable)} />,
  };

  // if (!(status?.items?.length > 0)) {
  //   return <>Loading...</>;
  // }

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={indexes.nItems}
      fetchFn={fetchUnchained}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={[indexes.updater, manifests.updater]} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
