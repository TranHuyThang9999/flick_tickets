import React, { useState } from "react";

import HomeaAdmin from "./dashboard/HomeaAdmin";
import CreateAccountStaff from "./dashboard/CreateAccountStaff";
import GetAllStaff from "./dashboard/GetAllStaff";
import SelectedSeat from "./common/cinemas/SelectedSeat";
import AppSquare from "./common/test/test";
import DisplaySelectedSeat from "./dashboard/DisplaySelectedSeat";
import GetTicketById from "./Home/Tickets/DetailTicketById";
import ListFileByTicketId from "./Home/Tickets/ListFileByTicketId";
import AdminUploadTickets from "./dashboard/AdminUploadTickets";
import TestDetailedShowSchedule from "./common/test/tesst1";
import SendEmail from "./common/ProcessEmail/SendEmail";
import VerifyEmail from "./common/ProcessEmail/VerifyEmail";
import FormCheckOtpEmail from "./common/ProcessEmail/FormCheckOtpEmail";
import DemoSend from "./common/test/test";
import Cookies from "js-cookie"; // Import thư viện js-cookie
import CinemasAdd from "./common/cinemas/CinemasAdd";
import MovieAdd from "./common/MovieTypes/MovieUpload";
import GetAllTicketsForAdmin from "./dashboard/GetAllTickets";
import TestGetAllSelect from "./common/test/test";
import FormLogin from "./dashboard/FormLogin";
import UseContextToken from "./Routers/UseContextToken";
import FormRegisterCustomer from "./common/customers/FormRegisterCustomer";
import UpdateProfile from "./common/customers/UpdateProfile";
import Profile from "./common/customers/Profile";
import PageForUser from "./Home/Page/PageForUser";
import GetListCart from "./cart/GetListCart";
import UpdateCartById from "./cart/UpdateCartById";
import QRScanner from "./QRScanner/QRScanner";
import PurchaseHistory from "./Home/Tickets/PurchaseHistory";
import AddShowTime from "./dashboard/AddShowTime";
import UpdateTicketById from "./dashboard/UpdateTicketById";
import UpdateShowTimeById from "./dashboard/UpdateShowTimeById";
import TestCheckBox from "./common/test/test";
import CheckLogin from "./Routers/CheckLogin";
import OrderStatistics from "./Orders/OrderStatistics";
import Blogs from './dashboard/Blogs';
import LoginWithEmail from "./common/customers/LoginWithEmail";
import UpSertFileByTicketId from "./dashboard/UpSertFileByTicketId";
import GetListTicketByMovieName from "./common/ViewTicketsForSell/GetListTicketByMovieName";
import TestSelect from "./common/test/test";
import RevenueOrder from "./Orders/RevenueOrder";
import StatisticalBar from "./Orders/StatisticalBar";

export default function App() {
  return (
    <div>
      {/* <HomeaAdmin /> */}
      {/* <CreateAccountStaff/> */}
      {/* <GetAllStaff/> */}
      {/* <AppTest/> */}
      {/* <SelectedSeat SelectedSeatGetFormApi={[1,12,13,14,15]} numSquares={100} heightContainerUseSaveData={500} widthContainerUseSavedate={900} onCreate={setSelectPopChid}/> */}
      {/* <AppSquare/> */}
      {/* <DisplaySelectedSeat/> */}
      {/* <GetTicketById id={7133536}/> */}
      {/* <GetListFileByTicketId id={7133536}/> */}
      {/* <PageForUser/> */}
      {/* <GetListTicket/> */}
      {/**/}
      {/* <ListFileByTicketId id={1485296}/> */}
      {/* <AdminUploadTickets/> */}
      {/* <TestDetailedShowSchedule/> */}
      {/* <SendEmail/>
      <VerifyEmail/> */}
      {/* <FormCheckOtpEmail/> */}
      {/* <DemoSend/> */}
      {/* <GetTicketById id={1485296} /> */}
      {/* <CinemasAdd/> */}
      {/* <MovieAdd/> */}
      {/* <GetAllTicketsForAdmin/> */}
      {/* <GetTicketByIdOnForm id={3244826}/> */}
      {/* <TestGetAllSelect/> */}
      {/* <FormLogin/> */}

      {/* <CheckLogin/> */}
      {/* <FormRegisterCustomer/> */}
      {/* <UpdateProfile/> */}
      {/* <Profile/> */}
      {/* <PageForUser/> */}
      {/* <TestLayout/> */}
      {/* <GetTicketById id={1485296}/> */}
      {/* <TestDisPlayTicket/> */}
      {/* <GetListCart/> */}
      {/* <QRScanner/> */}
      {/* <FormLogin/> */}
      {/* <PurchaseHistory email={'thuynguyen151387@gmail.com'}/> */}
      {/* <AddShowTime ticketId={3462449}/> */}
      {/* <UpdateTicketById ticketId={1357952}/> */}
      {/* <UpdateShowTimeById show_time_id={3169598}/> */}
      {/* <TestCheckBox/> */}
      {/* <PageForUser/> */}
      {/* <HomeaAdmin /> */}
      {/* <OrderStatistics/> */}
      {/* <Blogs/> */}
      {/* <LoginWithEmail/> */}
      {/* <PasswordRetrieval/> */}
      {/* <UpSertFileByTicketId ticketId={2528471}/> */}
      {/* <GetListTicketByMovieName movieName={'s'}/> */}
      {/* <TestSelect/> */}
      {/* <RevenueOrder/> */}
      <CheckLogin/>
      {/* <StatisticalBar/> */}
    </div>
  );
}
