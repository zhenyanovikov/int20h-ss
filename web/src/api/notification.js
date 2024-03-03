import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { enqueueSnackbar } from "notistack";
import { generatePath } from "react-router-dom";
import { useTranslation } from "react-i18next";
import {
  SEND_EMAIL_URL,
  EMAIL_TEMPLATES_KEY,
  EMAIL_TEMPLATES_URL,
} from "../constants/notification";
import apiClient from "./apiClient";

async function sendEmail(email) {
  const res = await apiClient.post(SEND_EMAIL_URL, email);
  return res.data;
}

async function getEmailTemplates() {
  const res = await apiClient.get(EMAIL_TEMPLATES_URL);
  return res.data;
}

async function getEmailTemplate(id) {
  const res = await apiClient.get(generatePath(EMAIL_TEMPLATES_URL, { id }));
  return res.data;
}

async function createEmailTemplate(template) {
  const res = await apiClient.post(EMAIL_TEMPLATES_URL, template);
  return res.data;
}

async function updateEmailTemplate(id, template) {
  const res = await apiClient.put(
    generatePath(EMAIL_TEMPLATES_URL, { id }),
    template
  );
  return res.data;
}

async function deleteEmailTemplate(id) {
  const res = await apiClient.delete(generatePath(EMAIL_TEMPLATES_URL, { id }));
  return res.data;
}

export function useSendEmail() {
  const { t } = useTranslation();

  return useMutation({
    mutationFn: sendEmail,
    onSuccess: () => {
      enqueueSnackbar(t("snackbars.sendEmail.success"), {
        variant: "success",
      });
    },
    onError: () => {
      enqueueSnackbar(t("snackbars.sendEmail.error"), {
        variant: "error",
      });
    },
  });
}

export function useGetEmailTemplates() {
  return useQuery({
    queryKey: [EMAIL_TEMPLATES_KEY],
    queryFn: getEmailTemplates,
  });
}

export function useGetEmailTemplate(id) {
  return useQuery({
    queryKey: [EMAIL_TEMPLATES_URL, id],
    queryFn: () => getEmailTemplate(id),
  });
}

export function useCreateEmailTemplate() {
  const { t } = useTranslation();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: createEmailTemplate,
    onSuccess: () => {
      enqueueSnackbar(t("snackbars.createEmailTemplate.success"), {
        variant: "success",
      });
      queryClient.invalidateQueries({ queryKey: [EMAIL_TEMPLATES_KEY] });
    },
    onError: () => {
      enqueueSnackbar(t("snackbars.createEmailTemplate.error"), {
        variant: "error",
      });
    },
  });
}

export function useUpdateEmailTemplate(id) {
  const { t } = useTranslation();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (template) => updateEmailTemplate(id, template),
    onSuccess: () => {
      enqueueSnackbar(t("snackbars.updateEmailTemplate.success"), {
        variant: "success",
      });
      queryClient.invalidateQueries({ queryKey: [EMAIL_TEMPLATES_URL] });
      queryClient.invalidateQueries({ queryKey: [EMAIL_TEMPLATES_URL, id] });
    },
    onError: () => {
      enqueueSnackbar(t("snackbars.updateEmailTemplate.error"), {
        variant: "error",
      });
    },
  });
}

export function useDeleteEmailTemplate() {
  const { t } = useTranslation();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: deleteEmailTemplate,
    onSuccess: () => {
      enqueueSnackbar(t("snackbars.deleteEmailTemplate.success"), {
        variant: "success",
      });
      queryClient.invalidateQueries({ queryKey: [EMAIL_TEMPLATES_URL] });
    },
    onError: () => {
      enqueueSnackbar(t("snackbars.deleteEmailTemplate.error"), {
        variant: "error",
      });
    },
  });
}
