import React from "react";
import NavBar from "./components/NavBar";
import { Form, Layout } from "antd";

const { Header, Content } = Layout;

function Shipping() {
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
      ></Content>
    </Layout>
  );
}

export default Shipping;
