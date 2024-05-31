import React, {useState } from "react";
import { Button, Divider } from "antd";
import "../App.css";
import ShipFromForm from "./ShipFromForm";
import ShipToForm from "./ShipToForm";
import PackageForm from "./PackageForm";
import Recommendation from "./Recommendation";

const Shipping = () => {
  const [showRecommendation, setShowRecommendation] = useState(false);

  const handleClick = () => {
    setShowRecommendation(true);
  };

  if (showRecommendation) {
    return <Recommendation />;
  }
  return (
    <div className="shipping-info-container">
      <h1 className='header-text'>Create a Shipment</h1>
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
