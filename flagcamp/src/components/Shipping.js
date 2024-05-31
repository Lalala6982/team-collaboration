import React, { useState } from "react";
import { Button, Divider, message } from "antd";
import "../App.css";
import ShipFromForm from "./ShipFromForm";
import ShipToForm from "./ShipToForm";
import PackageForm from "./PackageForm";
import Recommendation from "./Recommendation";
import { getShippingOptions } from "../utils";

const Shipping = () => {
  const [showRecommendation, setShowRecommendation] = useState(false);
  const [loading, setLoading] = useState(false);

  const handleClick = async (data) => {
    setLoading(true);
    setShowRecommendation(true);
    try {
      await getShippingOptions(data);
      message.success("Get Options successfully");
    } catch (error) {
      message.error(error.message);
    } finally {
      setLoading(false);
    }
  };

  if (showRecommendation) {
    return <Recommendation />;
  }
  return (
    <div className="shipping-info-container">
      <h1 className="header-text">Create a Shipment</h1>
      <Divider style={{ color: "gray" }} />
      <main className="form-container">
        <section>
          <section className="column">
            <ShipFromForm />
          </section>
          <section className="column">
            <ShipToForm />
          </section>
        </section>
        <section className="column">
          <PackageForm />
          <Button
            type="primary"
            htmlType="submit"
            className="package-form-submit-button"
            onClick={handleClick}
          >
            Continue
          </Button>
        </section>
      </main>
    </div>
  );
};

export default Shipping;
