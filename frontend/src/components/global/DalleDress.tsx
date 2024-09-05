import { useEffect, useState } from "react";
import { Image, Text, Loader, Center } from "@mantine/core";
import { GetDalle } from "@gocode/app/App";

export const DalleDress: React.FC = () => {
  const [content, setContent] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);
  const [isImage, setIsImage] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setLoading(true);
    GetDalle()
      .then((result) => {
        // If the result starts with "data:image", treat it as an image
        if (result.startsWith("data:image")) {
          setContent(result);
          setIsImage(true);
        } else {
          setContent(result);
          setIsImage(false);
        }
      })
      .catch((err) => {
        console.error("Failed to load content:", err);
        setError("Failed to load content. Please try again later.");
      })
      .finally(() => {
        setLoading(false);
      });
  }, []); // You can add dependencies here, such as [address] if needed

  if (loading) {
    return (
      <Center>
        <Loader />
      </Center>
    );
  }

  if (error) {
    return <Text color="red">{error}</Text>;
  }

  return isImage ? (
    <Image src={content as string} width={640} height={640} alt="Fetched content" />
  ) : (
    <Text>{content}</Text>
  );
};
