import { useQueryClient, useQuery, useMutation } from "@tanstack/react-query";
import { enqueueSnackbar } from "notistack";
import { useTranslation } from "react-i18next";
import { generatePath } from "react-router-dom";
import {
  FACULTIES_KEY,
  FACULTIES_URL,
  FACULTY_URL,
} from "../constants/faculty";
import apiClient from "./apiClient";

async function getFaculties() {
  const res = await apiClient.get(FACULTIES_URL);
  return res.data;
}

async function getFaculty(id) {
  const res = await apiClient.get(generatePath(FACULTY_URL, { id }));
  return res.data;
}

async function createFaculty(faculty) {
  const res = await apiClient.post(FACULTIES_URL, faculty);
  return res.data;
}

async function updateFaculty(id, faculty) {
  const res = await apiClient.put(generatePath(FACULTY_URL, { id }), faculty);
  return res.data;
}

async function deleteFaculty(id) {
  const res = await apiClient.delete(generatePath(FACULTY_URL, { id }));
  return res.data;
}

export function useGetFaculties() {
  return useQuery({
    queryKey: [FACULTIES_KEY],
    queryFn: getFaculties,
  });
}

export function useGetFaculty(id) {
  return useQuery({
    queryKey: [FACULTY_URL, id],
    queryFn: () => getFaculty(id),
  });
}

export function useCreateFaculty() {
  const { t } = useTranslation();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: createFaculty,
    onSuccess: () => {
      enqueueSnackbar(t("snackbars.createFaculty.success"), {
        variant: "success",
      });
      queryClient.invalidateQueries({ queryKey: [FACULTIES_KEY] });
    },
    onError: () => {
      enqueueSnackbar(t("snackbars.createFaculty.error"), { variant: "error" });
    },
  });
}

export function useUpdateFaculty(id) {
  const { t } = useTranslation();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (faculty) => updateFaculty(id, faculty),
    onSuccess: () => {
      enqueueSnackbar(t("snackbars.updateFaculty.success"), {
        variant: "success",
      });
      queryClient.invalidateQueries({ queryKey: [FACULTIES_KEY] });
      queryClient.invalidateQueries({ queryKey: [FACULTIES_KEY, id] });
    },
    onError: () => {
      enqueueSnackbar(t("snackbars.updateFaculty.error"), { variant: "error" });
    },
  });
}

export function useDeleteFaculty() {
  const { t } = useTranslation();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: deleteFaculty,
    onSuccess: () => {
      enqueueSnackbar(t("snackbars.deleteFaculty.success"), {
        variant: "success",
      });
      queryClient.invalidateQueries({ queryKey: [FACULTIES_KEY] });
    },
    onError: () => {
      enqueueSnackbar(t("snackbars.deleteFaculty.error"), { variant: "error" });
    },
  });
}
