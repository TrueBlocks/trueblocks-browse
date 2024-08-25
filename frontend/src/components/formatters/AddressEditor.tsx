import React, { useEffect, useState } from "react";
import { AddrToName } from "@gocode/app/App";
import { base } from "@gocode/models";
import { Formatter, FormatterProps, knownType, Popup, AddressPopup } from ".";

export enum EditorMode {
  All = "all",
  NameOnly = "name_only",
  AddressOnly = "address_only",
}

interface AddressEditorProps extends FormatterProps {
  mode?: EditorMode;
}

export const AddressEditor = ({ value, value2, className, size = "md", mode = EditorMode.All }: AddressEditorProps) => {
  const [line1, setLine1] = useState<string>("");
  const [line2, setLine2] = useState<string>("");

  const givenName = value2 as string;
  const givenAddress = value as unknown as string;

  useEffect(() => {
    const formatAddress = async () => {
      let address = value as string;
      if (!address || address == "0x0") {
        setLine1(givenName);
        setLine2("");
        return;
      }

      AddrToName(address as unknown as base.Address).then((knownName) => {
        if (knownName || givenName) {
          setLine1(knownName ? knownName : givenName);
          setLine2(value);
        } else {
          setLine1(value);
          setLine2("");
        }
      });
    };
    formatAddress();
  }, [value, givenName]);

  const line1Type: knownType = "address-line1";
  const line2Type: knownType = "address-line2";

  const editor = <AddressPopup address={value} name={line1} onSubmit={(newValue: string) => console.log(newValue)} />;
  return (
    <Popup editor={editor}>
      {/* {givenAddress} */}
      <Formatter className={className} size={size} type={line1Type} value={line1} />
      {line2 ? <Formatter className={className} size={size} type={line2Type} value={line2} /> : <></>}
    </Popup>
  );
};
