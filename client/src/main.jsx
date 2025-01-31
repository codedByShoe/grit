import { createRoot } from "react-dom/client";
import { createInertiaApp } from "@inertiajs/react";

createInertiaApp({
  resolve: (name) => require(`./pages/${name}.tsx`),
  setup({ el, App, props }) {
    createRoot(el).render(<App {...props} />);
  }
});
