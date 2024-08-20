import { knownTypes } from ".";
import classes from "./Formatter.module.css";

const debug = true;

export const getDebugColor = (type: knownTypes): string => {
  var ret: string = "";
  if (debug) {
    switch (type) {
      case "address-and-name":
        ret = classes.orange;
        break;
      case "address-address-only":
        ret = classes.orange;
        break;
      case "address-name-only":
        ret = classes.orange;
        break;
      case "boolean":
        break;
      case "check":
        break;
      case "error":
        break;
      case "ether":
        ret = classes.brown;
        break;
      case "timestamp":
        ret = classes.blue;
        break;
      case "bytes":
        ret = classes.green;
        break;
      case "float":
        ret = classes.red;
        break;
      case "int":
        ret = classes.blue;
        break;
      case "appearance":
        ret = classes.lightblue;
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
