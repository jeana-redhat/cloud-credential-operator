// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationMenuSpec) DeepCopyInto(out *ApplicationMenuSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationMenuSpec.
func (in *ApplicationMenuSpec) DeepCopy() *ApplicationMenuSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationMenuSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleCLIDownload) DeepCopyInto(out *ConsoleCLIDownload) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleCLIDownload.
func (in *ConsoleCLIDownload) DeepCopy() *ConsoleCLIDownload {
	if in == nil {
		return nil
	}
	out := new(ConsoleCLIDownload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleCLIDownload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleCLIDownloadList) DeepCopyInto(out *ConsoleCLIDownloadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConsoleCLIDownload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleCLIDownloadList.
func (in *ConsoleCLIDownloadList) DeepCopy() *ConsoleCLIDownloadList {
	if in == nil {
		return nil
	}
	out := new(ConsoleCLIDownloadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleCLIDownloadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleCLIDownloadSpec) DeepCopyInto(out *ConsoleCLIDownloadSpec) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]Link, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleCLIDownloadSpec.
func (in *ConsoleCLIDownloadSpec) DeepCopy() *ConsoleCLIDownloadSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleCLIDownloadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleExternalLogLink) DeepCopyInto(out *ConsoleExternalLogLink) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleExternalLogLink.
func (in *ConsoleExternalLogLink) DeepCopy() *ConsoleExternalLogLink {
	if in == nil {
		return nil
	}
	out := new(ConsoleExternalLogLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleExternalLogLink) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleExternalLogLinkSpec) DeepCopyInto(out *ConsoleExternalLogLinkSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleExternalLogLinkSpec.
func (in *ConsoleExternalLogLinkSpec) DeepCopy() *ConsoleExternalLogLinkSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleExternalLogLinkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleLink) DeepCopyInto(out *ConsoleLink) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleLink.
func (in *ConsoleLink) DeepCopy() *ConsoleLink {
	if in == nil {
		return nil
	}
	out := new(ConsoleLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleLink) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleLinkList) DeepCopyInto(out *ConsoleLinkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConsoleLink, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleLinkList.
func (in *ConsoleLinkList) DeepCopy() *ConsoleLinkList {
	if in == nil {
		return nil
	}
	out := new(ConsoleLinkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleLinkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleLinkSpec) DeepCopyInto(out *ConsoleLinkSpec) {
	*out = *in
	out.Link = in.Link
	if in.ApplicationMenu != nil {
		in, out := &in.ApplicationMenu, &out.ApplicationMenu
		*out = new(ApplicationMenuSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleLinkSpec.
func (in *ConsoleLinkSpec) DeepCopy() *ConsoleLinkSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleLinkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleNotification) DeepCopyInto(out *ConsoleNotification) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleNotification.
func (in *ConsoleNotification) DeepCopy() *ConsoleNotification {
	if in == nil {
		return nil
	}
	out := new(ConsoleNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleNotification) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleNotificationList) DeepCopyInto(out *ConsoleNotificationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConsoleNotification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleNotificationList.
func (in *ConsoleNotificationList) DeepCopy() *ConsoleNotificationList {
	if in == nil {
		return nil
	}
	out := new(ConsoleNotificationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleNotificationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleNotificationSpec) DeepCopyInto(out *ConsoleNotificationSpec) {
	*out = *in
	if in.Link != nil {
		in, out := &in.Link, &out.Link
		*out = new(Link)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleNotificationSpec.
func (in *ConsoleNotificationSpec) DeepCopy() *ConsoleNotificationSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleNotificationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Link) DeepCopyInto(out *Link) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Link.
func (in *Link) DeepCopy() *Link {
	if in == nil {
		return nil
	}
	out := new(Link)
	in.DeepCopyInto(out)
	return out
}
