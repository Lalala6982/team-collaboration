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
  const [authed, setAuthed] = useState();
  const [currentTab, setCurrentTab] = useState("1");
  const [showLogin, setShowLogin] = useState(false);
  const [showSignup, setShowSignup] = useState(false);
  const [activeKey, setActiveTabKey] = useState("1");

  useEffect(() => {
    const token = localStorage.getItem("authToken");
    setAuthed(!token);
  }, []);

  const handleLogOut = () => {
    localStorage.removeItem("authToken");
    setAuthed(false);
  };

  const toggleLoginModal = () => {
    setShowLogin(true);
  };

  const closeLoginModal = () => {
    setShowLogin(false);
  };

  const toggleSignupModal = () => {
    setShowSignup(true);
  };

  const closeSignupModal = () => {
    setShowSignup(false);
  };

  const userMenu = (
    <Menu>
      <Menu.Item key="history" onClick={toggleSignupModal}>
        SignUp
      </Menu.Item>
      <Menu.Item key="logout" onClick={toggleLoginModal}>
        Login
      </Menu.Item>
    </Menu>
  );

  const handleTabChange = (key) => {
    setCurrentTab(key);
  };

  const renderContent = (key) => {
    switch (key) {
      case "1":
        return <HomePage />;
      case "2":
        return (<Shipping />);
      case "3":
        return (<OrderHistory/>);
      case "4":
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
            <TabPane tab="Shipping" key="2"/>
            <TabPane tab ="Tracking" key = "3"/>
            <TabPane tab="About Us" key="4" />
          </Tabs>
          <div style={{
            marginLeft: 30, marginTop: 48
            
           }} alignItem="center">
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
      <Modal
        title="Login"
        visible={showLogin}
        onCancel={closeLoginModal}
        footer={null}
      >
        <LoginForm />
      </Modal>
      <Modal
        title="Sign Up"
        visible={showSignup}
        onCancel={closeSignupModal}
        footer={null}
      >
        <SignupButton />
      </Modal>
    </Layout>
  );
};

export default App;

