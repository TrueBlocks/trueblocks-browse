import React, { useEffect, useState } from "react";
import { TextProps } from "@mantine/core";
import { AddrToName } from "@gocode/app/App";
import { base } from "@gocode/models";
import { Formatter, FormatterProps, knownType, Popup, NamePopup, AddressPopup } from ".";
import { useAppState } from "@state";
import { types } from "@gocode/models";

export const AddressFormatter = ({ value, value2, className, size = "md" }: FormatterProps) => {
  const [line1, setLine1] = useState<string>("");
  const [line2, setLine2] = useState<string>("");

  useEffect(() => {
    const formatAddress = async () => {
      let address = value as string;
      if (!address || address == "0x0") {
        setLine1(value2);
        setLine2("");
        return;
      }

      AddrToName(address as unknown as base.Address).then((knownName) => {
        if (knownName || value2) {
          setLine1(knownName ? knownName : value);
          setLine2(value);
        } else {
          setLine1(value);
          setLine2("");
        }
      });
    };
    formatAddress();
  }, [value, value2]);

  const line1Type: knownType = "address-big";
  const line2Type: knownType = "address-small";

  const editor = <NamePopup address={value} name={line1} onSubmit={(newValue: string) => console.log(newValue)} />;
  return (
    <Popup editor={editor}>
      <Formatter className={className} size={size} type={line1Type} value={line1} />
      {line2 ? <Formatter className={className} size={size} type={line2Type} value={line2} /> : <></>}
    </Popup>
  );
};
