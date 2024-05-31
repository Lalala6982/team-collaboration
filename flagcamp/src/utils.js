const domain = "http://localhost:8080";

const handleResponseStatus = (response, errMsg) => {
  const { status, ok } = response;

  if (status === 401) {
    localStorage.removeItem("authToken"); // web storage
    window.location.reload();
    return;
  }

  if (!ok) {
    throw Error(errMsg);
  }
};

export const login = (credential) => {
  const url = `${domain}/signin`;
  return fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(credential),
  })
    .then((response) => {
      if (!response.ok) {
        throw Error("Fail to log in");
      }

      return response.text();
    })
    .then((token) => {
      localStorage.setItem("authToken", token);
    });
};

export const register = (credential) => {
  const url = `${domain}/signup`;
  return fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(credential),
  }).then((response) => {
    handleResponseStatus(response, "Fail to register");
  });
};

export const getShippingOptions = (data) => {
  const url = `${domain}/upload`;

  const {
    shipper,
    fromAdress,
    fromZipCode,
    fromCity,
    fromCounty,
    fromPhone,
    fromEmail,
    consignee,
    toAdress,
    toZipCode,
    toCity,
    toCounty,
    toPhone,
    toEmail,
    totalWeight,
  } = data;
    
  const formData = new FormData();
  formData.append("shipper", shipper);
  formData.append("from_address", fromAdress);
  formData.append("from_zip_code", fromZipCode);
  formData.append("from_city", fromCity);
  formData.append("from_county", fromCounty);
  formData.append("from_phone", fromPhone);
  formData.append("from_email", fromEmail);
  formData.append("consignee", consignee);
  formData.append("to_address", toAdress);
  formData.append("to_zip_code", toZipCode);
  formData.append("to_city", toCity);
  formData.append("to_county", toCounty);
  formData.append("to_phone", toPhone);
  formData.append("to_email", toEmail);
  formData.append("total_weight", totalWeight);

  return fetch(url, {
    method: "POST",
      headers: {
        "Content-Type": "application/json",
     },
    body: formData,
  }).then((response) => {
    handleResponseStatus(response, "Fail to upload app");
  });
};


export const createOrder = (data) => {
    const url = `${domain}/upload`;
 
    const {
      status,
      orderTime,
      productID,
      priceID,
      price,
      deiverID,
      duration,
      distance,
    } = data;
      
    const formData = new FormData();
    formData.append("status", status);
    formData.append("orderTime", orderTime);
    formData.append("productID", productID);
    formData.append("priceID", priceID);
    formData.append("price", price);
    formData.append("deiverID", deiverID);
    formData.append("duration", duration);
    formData.append("distance", distance);
  
    return fetch(url, {
      method: "POST",
        headers: {
          "Content-Type": "application/json",
       },
      body: formData,
    }).then((response) => {
      handleResponseStatus(response, "Fail to upload app");
    });
};
  
export const searchOrder = (query) => {
  const id = query?.id ?? "";
  const status = query?.status ?? "";

  const url = new URL(`${domain}/search`);
  url.searchParams.append("id", id);

  return fetch(url, {
    headers: {
      "Content-Type": "application/json",
    },
  }).then((response) => {
    handleResponseStatus(response, "Fail to search order");

    return response.json();
  });
};

export const getOrderHistory = (query) => {
  const id = query?.id ?? "";
  const orderTime = query?.orderTime ?? "";
  const status = query?.status ?? "";

  const authToken = localStorage.getItem("authToken");
  const url = new URL(`${domain}/search`);
  url.searchParams.append("id", id);
  url.searchParams.append("order_time", orderTime);
  url.searchParams.append("status", status);

  return fetch(url, {
    headers: {
      Authorization: `Bearer ${authToken}`,
    },
  }).then((response) => {
    handleResponseStatus(response, "Fail to get order history");

    return response.json();
  });
};

export const checkout = (orderId) => {
  const url = `${domain}/checkout?optionID=${orderId}`;

  return fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => {
      handleResponseStatus(response, "Fail to checkout");

      return response.text();
    })
    .then((redirectUrl) => {
      window.location = redirectUrl;
    });
};
