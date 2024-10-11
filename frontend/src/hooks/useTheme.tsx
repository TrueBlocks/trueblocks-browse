import { useMantineTheme } from "@mantine/core";

interface UseThemeProps {
  which: "background" | "text";
}

export const useTheme = ({ which }: UseThemeProps) => {
  const theme = useMantineTheme();
  switch (which) {
    case "background":
      return theme.colors[theme.primaryColor][6];
  }
};
