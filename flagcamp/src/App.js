import React from "react";
import NavBar from "./components/NavBar";
import Tracking from "./components/Tracking";
import { Row, Layout } from "antd";
import "./App.css";
import "./index.css"

const { Header, Content } = Layout;

function App() {
  return (
    <Layout style={{ height: "100vh" }}>
      <Header
        className="site-header-backgroud"
        style={{ display: "flex", justifyContent: "space-between" }}
      >
        <div className="site-name-font" >Shipping Service</div>
        <NavBar />
      </Header>
      <Content
        className="site-layout-background"
        style={{
          padding: 300,
          height: "calc(100% - 64px)",
          overflow: "auto",
        }}
      >
        <Tracking />
      </Content>
    </Layout>
  );
}

export default App;
