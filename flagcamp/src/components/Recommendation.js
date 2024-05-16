import React from "react";
import { Form, Divider, Button } from "antd";
import ShippingOptions from "./ShippingOptions";
import OrderSummary from "./OrderSummary";

const { Item } = Form;

const Recommendation = () => {
  const handleClick = () => {};
  return (
    <div className="shipping-info-container">
      <h1 style={{ fontSize: 40 }}>Select a Shipping Service Option</h1>
      <Divider style={{ color: "gray" }} />
      <main className="form-container">
        <section>
          <section className="column">
            <ShippingOptions />
          </section>
        </section>
        <section className="column">
          <OrderSummary/>
          <Button
            type="primary"
            htmlType="submit"
            className="package-form-submit-button"
            style={{borderRadius: 5}}
            onClick={handleClick}
          >
            Continue to payment
          </Button>
        </section>
      </main>
    </div>
  );
};

export default Recommendation;
