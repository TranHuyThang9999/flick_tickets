import React from 'react'
import AdminUploadTickets from './dashboard/AdminUploadTickets'
import CinemasAdd from './common/cinemas/CinemasAdd'
import HomeaAdmin from './dashboard/HomeaAdmin'
import DisplayTickets from './common/customers/DisplayTickets'
import QRScanner from './QRScanner/QRScanner'
import CheckQr from './QRScanner/CheckQr'

export default function App() {
  const token = 'mz20CA3h%2Fn0ybLeU0uMJCg%3D%3D';

  return (
    <div>
      {/* <AdminUploadTickets/> */}
      {/* <CinemasAdd/> */}
      {/* <HomeaAdmin/> */}
      {/* <DisplayTickets/> */}
      <QRScanner/>
      {/* <CheckQr token={token}/> */}
   </div>
  )
}
