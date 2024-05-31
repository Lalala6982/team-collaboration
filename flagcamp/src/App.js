import React, { useEffect, useState } from "react";
import { Layout, Dropdown, Menu, Button, Tabs, Modal } from "antd";
import HomePage from "./components/HomePage";
import About from "./components/About";
import Shipping from "./components/Shipping";
import { UserOutlined } from "@ant-design/icons";
import SignupButton from "./components/SignupButton";
import LoginForm from "./components/LoginForm";
import OrderSummary from "./components/OrderSummary";
import OrderHistory from "./components/OrderHistory";

const { Header, Content } = Layout;
const { TabPane } = Tabs;

const App = () => {
  const [authed, setAuthed] = useState(); // Remember to set to false at final implementation
  const [currentTab, setCurrentTab] = useState("1");
  const [activeKey, setActiveTabKey] = useState("1");

  useEffect(() => {
    const authToken = localStorage.getItem("authToken");
    setAuthed(authToken !== null);
  }, []);
  const handleLoginSuccess = () => {
    setAuthed(true);
  };
  const handleLogOut = () => {
    localStorage.removeItem("authToken");
    setAuthed(false);
  };

  const userMenu = () => {
    if (authed) {
      return (
        <Menu>
          <Menu.Item key="logout" onClick={handleLogOut}>
            LogOut
          </Menu.Item>
        </Menu>
      );
    }
    return (
      <Menu>
        <Menu.Item key="signup">
          <SignupButton />
        </Menu.Item>
        <Menu.Item key="login">
          <LoginForm onLoginSuccess={handleLoginSuccess} />
        </Menu.Item>
      </Menu>
    );
  };

  const handleTabChange = (key) => {
    setCurrentTab(key);
  };
  const renderContent = (key) => {
    switch (key) {
      case "1":
        return <HomePage authed={authed} />;
      case "2":
        return <Shipping />;
      case "3":
        return <About />;
    }
  };

  return (
    <Layout style={{ height: "100vh" }}>
      <Header
        className="site-header-backgroud"
        style={{ display: "flex", justifyContent: "space-between" }}
      >
        <div className="site-name-font">Shipping Service</div>
        <div style={{ display: "flex", justifyContent: "space-between" }}>
          <Tabs
            defaultActiveKey="1"
            onChange={handleTabChange}
            destroyInactiveTabPane={true}
            className="equal-width-tabs"
          >
            <TabPane tab="Home" key="1" />
            <TabPane tab="Shipping" key="2" />
            <TabPane tab="About Us" key="3" />
          </Tabs>
          <div
            style={{
              marginLeft: 30,
              marginTop: 48,
            }}
            alignItem="center"
          >
            <Dropdown trigger="click" overlay={userMenu}>
              <Button icon={<UserOutlined />} shape="circle" />
            </Dropdown>
          </div>
        </div>
      </Header>
      <Content
        className="site-layout-background"
        style={{
          paddingTop: 20,
          paddingLeft: 100,
          paddingRight: 100,
          height: "calc(100% - 64px)",
          overflow: "auto",
        }}
      >
        {renderContent(currentTab)}
      </Content>
    </Layout>
  );
};

export default App;
