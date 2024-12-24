import React from "react";
import AppHead from "@/components/AppHead";
import Navbar from "@/components/Navbar";

const AppLayout = ({ title, children }) => {
  return (
    <div>
      <Navbar />
      <main>
        <AppHead title={title} />
        <div className="mx-auto max-w-7xl sm:px-6 lg:px-8 py-12">
          {/* Page Content */}
          {children}
        </div>
      </main>
    </div>
  )
}

export default AppLayout
