import { CellType } from "@components";
import { useEnvironment } from "@hooks";
import classes from "./Formatter.module.css";

export const GetDebugColor = (type: CellType): string => {
  const debug = useEnvironment("TB_DEBUG_DISPLAY");

  let ret: string = "";
  if (debug) {
    switch (type) {
      case "address-editor":
      case "address-address-only":
      case "address-name-only":
      case "address-line1":
      case "address-line2":
        ret = classes.orange;
        break;
      case "boolean":
        break;
      case "check":
        break;
      case "error":
        break;
      case "tag":
        ret = classes.green;
        break;
      case "ether":
        ret = classes.red;
        break;
      case "timestamp":
        ret = classes.red;
        break;
      case "bytes":
        ret = classes.red;
        break;
      case "float":
        ret = classes.red;
        break;
      case "int":
        ret = classes.red;
        break;
      case "appearance":
        ret = classes.lightblue;
        break;
      case "time":
        ret = classes.purple;
        break;
      case "date":
        ret = classes.purple;
        break;
      case "hash":
        ret = classes.lightblue;
        break;
      case "path":
        ret = classes.pink;
        break;
      case "range":
        ret = classes.pink;
        break;
      case "text":
        ret = classes.green;
        break;
      case "url":
        ret = classes.red;
        break;
      default:
        ret = classes.black;
        break;
    }
  }
  return ret;
};
