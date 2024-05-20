import React, { useState } from 'react';
import { Table, Input, Button, Space, Divider, Row, Col } from 'antd';
import { SearchOutlined } from '@ant-design/icons';

const data = [
  { key: '1', date: 'Dec 5', orderId: 123486, trackingId: 5098021, status: 'Delivered' },
  { key: '2', date: 'Dec 5', orderId: 650890, trackingId: 5098021, status: 'Delivered' },
  { key: '3', date: 'Dec 5', orderId: 890566, trackingId: 5098021, status: 'Delivered' },
  { key: '4', date: 'Dec 5', orderId: 210400, trackingId: 5098021, status: 'Delivered' },
];

function OrderHistory() {
  const [searchText, setSearchText] = useState('');
  
  const handleSearch = (selectedKeys, confirm) => {
    confirm();
    setSearchText(selectedKeys[0]);
  };

  const handleReset = (clearFilters) => {
    clearFilters();
    setSearchText('');
  };

  const columns = [
    {
      title: 'Date',
      dataIndex: 'date',
      key: 'date',
    },
    {
      title: 'Order ID',
      dataIndex: 'orderId',
      key: 'orderId',
    },
    {
      title: 'Tracking ID',
      dataIndex: 'trackingId',
      key: 'trackingId',
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status',
    }
  ];

  return(
    <div>
        <Row justify="space-between">
            <Col>
                <h1 style={{ fontSize: 40 }}>Order History</h1>
            </Col>
            <Col>
                <Space>
                    <Input placeholder="Search tickets..." prefix={<SearchOutlined />}/>
                    <Button  type="primary">Filter</Button>
                </Space>
                
            </Col>
        </Row>
        <Divider/>
        <Table columns={columns} dataSource={data} />
    </div>
  ) 
}

export default OrderHistory;
