import React, { useState, useEffect } from "react";
import { Table, Input, Button, Space, Divider, Row, Col } from "antd";
import { SearchOutlined } from "@ant-design/icons";
import "../App.css";
import { getOrderHistory } from "../utils";

function OrderHistory() {
  const [orderId, setOrderId] = useState("");
  const [orderData, setOrderData] = useState([]);
  const [loading, setLoading] = useState(false);

  const fetchOrderHistory = async (query = {}) => {
    setLoading(true);
    try {
      const data = await getOrderHistory();
      console.log("data", data)
      setOrderData(data);
    } catch (error) {
      console.error("Error fetching order history:", error);
    }
    setLoading(false);
  };

  useEffect(() => {
    fetchOrderHistory();
  }, []);

  const handleSearch = () => {
    // fetchOrderHistory({ id: orderId });
  };

  const columns = [
    {
      title: "Date",
      dataIndex: "order_time",
      key: "date",
    },
    {
      title: "Order ID",
      dataIndex: "id",
      key: "orderId",
    },
    {
      title: "Tracking ID",
      dataIndex: "id",
      key: "trackingId",
    },
    {
      title: "Status",
      dataIndex: "status",
      key: "status",
    },
  ];

  return (
    <div>
      <Row justify="space-between">
        <Col>
          <h1 className="header-text">Order History</h1>
        </Col>
        <Col>
          <Space>
            <Input
              placeholder="Order ID"
              value={orderId}
              onChange={(e) => setOrderId(e.target.value)}
              prefix={<SearchOutlined />}
            />
            <Button type="primary" onClick={handleSearch} loading={loading}>
              Filter
            </Button>
          </Space>
        </Col>
      </Row>
      <Divider />
      <Table columns={columns} dataSource={orderData} loading={loading} />
    </div>
  );
}

export default OrderHistory;

