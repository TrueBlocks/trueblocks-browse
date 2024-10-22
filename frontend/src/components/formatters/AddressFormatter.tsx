import { useEffect, useState, useCallback } from "react";
import { Formatter, FormatterProps, CellType, Popup, AddressPopup } from "@components";
import { AddrToName, ModifyName } from "@gocode/app/App";
import { app } from "@gocode/models";
import { useAppState, useViewState } from "@state";
import classes from "./Formatter.module.css";

export enum EdMode {
  All = "all",
  Name = "name_only",
  Address = "address_only",
}

interface AddressEditorProps extends FormatterProps {
  mode?: EdMode;
}

export const AddressFormatter = ({ value, value2, className, mode = EdMode.All }: Omit<AddressEditorProps, "size">) => {
  const { address, fetchNames } = useAppState();
  const { pager } = useViewState();

  const [line1, setLine1] = useState<string>("");
  const [line2, setLine2] = useState<string>("");
  const [isPopupOpen, setPopupOpen] = useState(false);

  const givenName = value2 as string;
  const givenAddress = value as unknown as string;

  useEffect(() => {
    const formatAddress = async () => {
      if (!givenAddress || givenAddress === "0x0") {
        setLine1(givenName);
        setLine2("");
        return;
      }

      switch (mode) {
        case EdMode.Address:
          setLine1(givenAddress);
          break;
        case EdMode.Name:
          AddrToName(value).then((knownName) => {
            if (knownName || givenName) {
              setLine1(knownName ? knownName : givenName);
              setLine2("");
            } else {
              setLine1("");
              setLine2("");
            }
          });
          break;
        case EdMode.All:
        default:
          AddrToName(value).then((knownName) => {
            if (knownName || givenName) {
              setLine1(knownName ? knownName : givenName);
              setLine2(value);
            } else {
              setLine1(value);
              setLine2("");
            }
          });
      }
    };
    formatAddress();
  }, [value, value2, mode]);

  const line1Type: CellType = "address-line1";
  const line2Type: CellType = "address-line2";

  const onClose = useCallback(() => setPopupOpen(false), []);

  const editor = isPopupOpen ? (
    <AddressPopup
      address={value}
      name={line1}
      onClose={onClose}
      onSubmit={(newValue: string) => {
        setPopupOpen(false);
        const modData = app.ModifyData.createFrom({
          operation: "update",
          address: givenAddress,
          value: newValue,
        });
        ModifyName(modData).then(() => {
          fetchNames(pager.getOffset(), pager.perPage);
        });
      }}
    />
  ) : null;

  const isCurrent = value === address;
  const line1Class = isCurrent ? classes.bold : className;

  return (
    <Popup editor={editor}>
      <div onClick={() => setPopupOpen(true)}>
        <Formatter className={line1Class} type={line1Type} value={line1} />
        {line2 ? <Formatter className={className} type={line2Type} value={line2} /> : null}
      </div>
    </Popup>
  );
};
