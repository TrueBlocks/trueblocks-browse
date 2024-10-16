import React from "react";

export function validAddress(value: string) {
  if (!/^0x[a-fA-F0-9]{40}$/.test(value) && !/^[a-zA-Z0-9.-]+\.eth$/.test(value)) {
    return "Please enter a valid Ethereum address.";
  }
  return null;
}

export function requiredString(value: string) {
  if (value.length === 0) {
    return "Please enter a non-empty value.";
  }
  return null;
}
