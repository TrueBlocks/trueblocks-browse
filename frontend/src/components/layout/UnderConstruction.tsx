import { Text } from "@mantine/core";

type UnderConstructionProps = {
  message: string;
};

export const UnderConstruction = ({ message }: UnderConstructionProps) => {
  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        width: "100%",
        height: "100%",
        position: "relative",
      }}
    >
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <Text
        style={{
          fontSize: "2rem", // Larger text size
          fontWeight: "bold", // Bold text
          position: "absolute",
          top: "33%", // Positioned one-third of the way down
          transform: "translateY(-50%)", // Ensure proper centering
          textAlign: "center",
        }}
      >
        {message}
      </Text>
    </div>
  );
};
