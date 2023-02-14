import Navbar from "./Navbar";
import ChangePassword from "./ChangePassword";
import NavigateDetailProperty from "./NavbarDetailProperty";
// import HomeLogin from "../pages/HomeLogin";
import { Button, Container, Row, Col, Form, Image } from "react-bootstrap";
import React, { useContext, useEffect } from "react";
// IMPORT IMAGES
import Name from "../assets/images/IP.png";
import Ema from "../assets/images/Email.png";
import Phone from "../assets/images/Phone.png";
import Place from "../assets/images/Tendant.png";
import Robin from "../assets/images/wowo.jpeg";
import Ad from "../assets/images/Lock.jpg";
import G from "../assets/images/Gender.png";
import Password from "../assets/images/Pass.png";
import { UserContext } from "../context/userContext";
import { useState } from "react";
// import home from "../pages/Home";
// import { Pass } from "react-bootstrap-icons";
import robin from "../assets/images/wowo.jpeg"

import { useQuery } from "react-query";
import { API } from "../config/api";
import { useMutation } from "react-query";
import jwt from "jwt-decode";

export default function Profile(props) {
  // useEffect(() => {
  //   document.body.style.background = "rgba(196, 196, 196, 0.25)";
  // });

  const [modalShow, setModalShow] = React.useState(false);

  const getToken = localStorage.getItem("token");

  const hasilDecode = jwt(getToken);

  // Fetching data from database
  const { data: user, refetch } = useQuery("userCache", async () => {
    const response = await API.get("/user/" + hasilDecode.id);
    console.log("ini response nyaa", response);
    return response.data.data;
  });

  console.log("iniiii", user)

  const [photo, setPhoto] = useState({
    image: "",
  });

  const handleChangePhoto = (e) => {
    const { name, type } = e.target;
    setPhoto({
      ...photo,
      [name]: type === "file" ? e.target.files : e.target.value,
    });

    if (e.target.type === "file") {
      let url = URL.createObjectURL(e.target.files[0]);
      console.log("blob image", url);
    }
  };

  const handleUpdate = useMutation(async (e) => {
    try {
      e.preventDefault();

      const formData = new FormData();
      if (photo.image) {
        formData.set("image", photo?.image[0], photo?.image[0]?.name);
      }

      const response = await API.patch("/user/" + hasilDecode.id, formData);
      console.log("berhasil mengubah photo profile", response);

      photo.image = "";
      refetch();
    } catch (error) {
      console.log(error);
    }
  });

  return (
    <div className="" style={{ marginTop: "9rem" }}>
      {/* <HomeLogin /> */}
      <Navbar
      />
      <Container className="">
        <Row className="justify-content-between bg-white py-3">
          <Col className="d-flex flex-column gap-3" sm={4}>
            <h3 className="fw-bold ">Personal Info</h3>
            <div className="d-flex align-items-center gap-3">
              <div>
                <img width={40} src={Name} alt="" />
              </div>
              <div className="d-flex flex-column">
                <span className="p-0 m-0 fw-bold">{user?.fullname}</span>
                <span className="fs14 text-secondary">Full Name</span>
              </div>
            </div>
            <div className="d-flex align-items-center gap-3">
              <img width={40} src={Ema} alt="" />
              <div className="d-flex flex-column">
                <span className="p-0 m-0 fw-bold">{user?.email}</span>
                <span className="fs14 text-secondary">Email</span>
              </div>
            </div>
            <div className="d-flex align-items-center gap-3">
              <img width={40} src={Password} alt="" />
              <div className="d-flex align-items-center gap-3">
                <div className="d-flex flex-column">
                  <Button
                    onClick={() => setModalShow(true)}
                    className="btn btn-dark bg-white text-primary fw-bold p-0 m-0 border-0"
                  >
                    Change Password
                  </Button>
                  <ChangePassword
                    show={modalShow}
                    onHide={() => setModalShow(false)}
                  />
                  <span className="p-0 m-0">Password</span>
                </div>
              </div>
            </div>
            {/* <div className="d-flex flex-column">
                <span className="p-0 m-0" fw-bold>Change Password</span>
                <span className="fs14 text-secondary">Password</span>
              </div> */}
            <div className="d-flex align-items-center gap-3">
              <img width={39} src={Place} alt="" />
              <div className="d-flex flex-column">
                <span className="p-0 m-0 fw-bold">
                  {user?.listAsRole}
                </span>
                <span className="fs14 text-secondary">Status</span>
              </div>
            </div>
            <div className="d-flex align-items-center gap-3">
              <img width={39} src={G} alt="" />
              <div className="d-flex flex-column">
                <span className="p-0 m-0 fw-bold">{user?.gender}</span>
                <span className="fs14 text-secondary">Gender</span>
              </div>
            </div>
            <div className="d-flex align-items-center gap-3">
              <img width={40} src={Phone} alt="" />
              <div className="d-flex flex-column">
                <span className="p-0 m-0 fw-bold">{user?.phone}</span>
                <span className="fs14 text-secondary">Mobile phone</span>
              </div>
            </div>
            <div className="d-flex align-items-center gap-3">
              <img width={38} src={Ad} alt="" />
              <div className="d-flex flex-column">
                <span className="p-0 m-0 fw-bold">{user?.address}</span>
                <span className="fs14 text-secondary">Address</span>
              </div>
            </div>
          </Col>
          <Col className=" d-flex flex-column gap-3 px-3 me-2" sm={4}>
            <div>
              <Image className="w-100 rounded pt-4" src={robin} />
            </div>
            <div>
              <Button className="w-100">Change Foto Profile</Button>
            </div>
          </Col>
          {/* <Col>
          
              <Form
                onSubmit={(e) => handleUpdate.mutate(e)}
                className="mt-2 w-100"
                style={{
                  maxWidth: "18rem",
                  objectFit: "cover",
                }}
              >
                <Form.Group className="mb-5">
                  <Form.Control
                    type="file"
                    id="upload"
                    name="image"
                    onChange={handleChangePhoto}
                    required
                  />
                  <Button
                    className="btn btn-primary mt-2 w-100"
                    type="submit"
                    style={{
                      maxWidth: "18rem",
                      objectFit: "cover",
                    }}
                  >
                    Change Photo Profile
                  </Button>
                </Form.Group>
              </Form>

          </Col> */}
        </Row>
      </Container>
    </div>
  );
}
