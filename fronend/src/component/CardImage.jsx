import React from "react";
import { getApi } from "../api/useApi.js";

// eslint-disable-next-line react/prop-types
function CardImage({ url }) {
  const [images, setImage] = React.useState();

  React.useEffect(() => {
    async function getImage() {
      const { response, error } = await getApi(`public/${url}`,null, {
        responseType: `blob`,
      });
      if (response) {
        const objUrl = URL.createObjectURL(response);
        setImage(objUrl);
      }
      if (error) {
        return null;
      }
    }

    getImage();

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return <img src={images} alt="ERROR BRO SORRY MY BAD" />;
}

export default CardImage;
