import React from "react";

import { Head } from '@inertiajs/react'

const AppHead = ({ title, children }) => {
  const pageTitle = (title) => { return `${title} | Go React` }
  return (
    <Head>
      <title>{title ? pageTitle(title) : pageTitle('Welcome')}</title>
      {children}
    </Head>
  )
}

export default AppHead
