import { createInertiaApp } from "@inertiajs/react";
import { createRoot } from "react-dom/client";
import React from "react";
import { ThemeProvider } from "@/components/ThemeProvider";

createInertiaApp({
  progress: {
    delay: 250,
    color: "#29d",
    includeCSS: true,
    showSpinner: true,
  },
  resolve: async (name) => {
    try {
      const module = await import(`./pages/${name}.jsx`);
      return module.default;
    } catch (err) {
      console.error(`Failed to load page: ${name}`, err);
      return null;
    }
  },
  setup({ el, App, props }) {
    createRoot(el).render(
      <ThemeProvider>
        <App {...props} />
      </ThemeProvider>,
    );
  },
});
