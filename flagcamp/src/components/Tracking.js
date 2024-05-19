import React, { useState } from "react";
import { Row, Button, Input, Form, message, List, Card } from "antd";

function Tracking() {
  const [loading, setLoading] = useState(false);
  const [trackingID, setTrackingID] = useState("");
  const [trackingStatus, setTrackingStatus] = useState([]);
  const [isVisible, setIsVisible] = useState(true);

  const handleTrack = () => {
    if (trackingID === "") {
      return;
    }
    setLoading(true);
    setIsVisible(true);
    // try {
    //   const data = await getTrackingStatus(trackingID);
    //   setTrackingStatus(data.result);
    // } catch (error) {
    //   message.error(error.message);
    // } finally {
    //   setLoading(false);
    //   setIsVisible(false;)
    // }
  };

  return (
    <Form style={{ maxWidth: 600 }} layout="vertical">
      <Form.Item
        style={{ padding: 2 }}
        label="Tracking ID"
        name="trackingid"
        rules={[
          {
            required: true,
            message: "Please enter tracking ID!",
          },
        ]}
      >
        <Input.Group compact>
          <Input
            style={{ width: 450, marginRight: 5 }}
            placeholder="Enter your Tracking ID"
            value={trackingID}
            onChange={(e) => setTrackingID(e.target.value)}
            onPressEnter={handleTrack}
          />
          <Button type="primary" onClick={handleTrack}>
            Track
          </Button>
        </Input.Group>
        <List
          loading={loading}
          style={{
            marginTop: 20,
            height: "calc(100% - 30px)",
            overflow: "auto",
            display: isVisible ? "block" : "none",
          }}
          grid={{ gutter: 1, xs: 1, sm: 1, md: 1, lg: 1, xl: 1, xxl: 1 }}
          dataSource={[1]}
          renderItem={(currentStatus) => (
            <List.Item key={currentStatus}>
              <Card title={currentStatus} />
            </List.Item>
          )}
        />
      </Form.Item>
    </Form>
  );
}
export default Tracking;
