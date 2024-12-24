import React from "react";
import AppHead from "../components/AppHead";
import { ModeToggle } from "../components/ModeToggle";

const GuestLayout = ({ title, children }) => {
  return (
    <>
      <AppHead title={title} />
      <ModeToggle />
      <div className="flex min-h-screen items-center justify-center">
        {/* Page Content */}
        {children}
      </div>
    </>
  )
}

export default GuestLayout
