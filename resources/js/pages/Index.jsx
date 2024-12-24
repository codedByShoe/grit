import React from "react";
import AppLayout from "@/layouts/AppLayout.jsx"
import { Button } from "@/components/ui/button.jsx"

const Index = () => {
  return (
    <AppLayout title={"Home"}>
      <Button>click</Button>
      <h1>Hello from react!</h1>
    </AppLayout>
  )
}

export default Index
