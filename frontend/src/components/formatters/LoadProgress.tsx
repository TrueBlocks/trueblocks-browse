import { useState, useEffect } from "react";
import { Progress, Text } from "@mantine/core";
import { FormatterProps } from "@components";

export const LoadProgress = ({ value, value2 }: Omit<FormatterProps, "type">) => {
  const [loaded, setLoaded] = useState<number>(0);
  const [total, setTotal] = useState<number>(0);

  useEffect(() => {
    setTotal(value2 as number);
  }, [value2]);

  useEffect(() => {
    setLoaded(Math.round((value * total) / 100));
  }, [value, total]);

  if (loaded === 0) {
    return (
      <>
        <Progress value={value} color="yellow" />
        <Text size="xs" mt="xs">
          not loaded
        </Text>
      </>
    );
  }

  if (value > 98) {
    return (
      <>
        <Progress value={value} color="green" />
        <Text size="xs" mt="xs">
          {loaded} of {total}
        </Text>
      </>
    );
  }

  return (
    <>
      <Progress value={value} color="yellow" striped />
      <Text size="xs" mt="xs">
        {loaded} of {total}
      </Text>
    </>
  );
};
