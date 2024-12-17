// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { useCallback } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { DebugState, TabItem, View, ViewForm } from "@components";
import { useNoops } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { IndexesFormDef, IndexesTableDef } from "../Indexes";
import { ManifestsFormDef, ManifestsTableDef } from "../Manifests";
import { PinsFormDef, PinsTableDef } from "../Pins";
import { PublishFormDef, PublishTableDef } from "../Publish";

// EXISTING_CODE

export const UnchainedView = () => {
  const { indexes, fetchIndexes, manifests, fetchManifests, pins, fetchPins, publish, fetchPublish } = useAppState();
  const { enterNoop, modifyNoop } = useNoops();
  const handleEnter = enterNoop;
  const handleModify = modifyNoop;

  const fetchUnchained = useCallback(
    (currentItem: number, itemsPerPage: number) => {
      fetchIndexes(currentItem, itemsPerPage);
      fetchManifests(currentItem, itemsPerPage);
      fetchPins(currentItem, itemsPerPage);
      fetchPublish(currentItem, itemsPerPage);
    },
    [fetchIndexes, fetchManifests, fetchPins, fetchPublish]
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

  const pinsTable = useReactTable({
    data: pins?.items || [],
    columns: PinsTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const publishTable = useReactTable({
    data: publish?.items || [],
    columns: PublishTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const tabItems: ViewForm = {
    indexes: <TabItem data={indexes} groups={IndexesFormDef(indexesTable)} />,
    manifests: <TabItem data={manifests} groups={ManifestsFormDef(manifestsTable)} />,
    pins: <TabItem data={pins} groups={PinsFormDef(pinsTable)} />,
    publish: <TabItem data={publish} groups={PublishFormDef(publishTable)} />,
  };

  // if (!(status?.items?.length > 0)) {
  //   return <>Loading...</>;
  // }

  return (
    <ViewStateProvider
      // do not remove - delint
      nItems={indexes.nItems}
      fetchFn={fetchUnchained}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={[indexes.updater, manifests.updater, pins.updater, publish.updater]} />
      <View tabItems={tabItems} />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
