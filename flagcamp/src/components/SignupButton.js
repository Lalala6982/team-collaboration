import { Button, Form, Modal, Input, message } from "antd";
import { useState } from "react";
import { UserOutlined } from "@ant-design/icons";
import { register } from "../utils";

const SignupButton = () => {
  const [modalVisible, setModalVisible] = useState(false);
  const [loading, setLoading] = useState(false);

  const handleRegisterOnClick = () => {
    setModalVisible(true);
  };
    
    const handleRegisterCancel = () => { 
        setModalVisible(false);
    }

  const handleFormSubmit = async (data) => {
    setLoading(true);

    try {
      //  await register(data);
      message.success("Sign up successfully!");
      setModalVisible(false);
    } catch (error) {
      message.error(error.message);
    } finally {
      setLoading(false);
    }
  };
    
  return (
    <div>
      <Form onFinish={handleFormSubmit}>
        <Form.Item
          name="username"
          rules={[{ required: true, message: "Please enter username!" }]}
        >
          <Input
            disabled={loading}
            prefix={<UserOutlined />}
            placeholder="Username"
          />
        </Form.Item>
        <Form.Item
          rules={[
            {
              required: true,
              message: "Please enter password!",
            },
          ]}
        >
          <Input
            disabled={loading}
            prefix={<UserOutlined />}
            placeholder="Password"
          />
        </Form.Item>
        <Form.Item>
          <Button
            loading={loading}
            type="primary"
            htmlType="submit"
            style={{ width: "100%" }}
            onClick={handleRegisterOnClick}
          >
            Sign Up
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};
export default SignupButton;
