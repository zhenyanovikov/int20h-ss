/* eslint-disable import/no-named-as-default-member */

import i18n from "i18next";
import { initReactI18next } from "react-i18next";
import enLocale from "../locales/en.json";
import ukLocale from "../locales/uk.json";
import { LOCALE } from "../constants/localisation";

i18n.use(initReactI18next).init({
  resources: {
    en: {
      translation: enLocale,
    },
    uk: {
      translation: ukLocale,
    },
  },
  lng: LOCALE.UK,
  fallbackLng: LOCALE.UK,
  interpolation: {
    escapeValue: false,
  },
});
