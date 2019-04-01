// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HangoutsChatNotifier) DeepCopyInto(out *HangoutsChatNotifier) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HangoutsChatNotifier.
func (in *HangoutsChatNotifier) DeepCopy() *HangoutsChatNotifier {
	if in == nil {
		return nil
	}
	out := new(HangoutsChatNotifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Notifier) DeepCopyInto(out *Notifier) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Notifier.
func (in *Notifier) DeepCopy() *Notifier {
	if in == nil {
		return nil
	}
	out := new(Notifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Notifier) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotifierList) DeepCopyInto(out *NotifierList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Notifier, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotifierList.
func (in *NotifierList) DeepCopy() *NotifierList {
	if in == nil {
		return nil
	}
	out := new(NotifierList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NotifierList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotifierSpec) DeepCopyInto(out *NotifierSpec) {
	*out = *in
	if in.Slack != nil {
		in, out := &in.Slack, &out.Slack
		*out = new(SlackNotifier)
		**out = **in
	}
	if in.HangoutsChat != nil {
		in, out := &in.HangoutsChat, &out.HangoutsChat
		*out = new(HangoutsChatNotifier)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotifierSpec.
func (in *NotifierSpec) DeepCopy() *NotifierSpec {
	if in == nil {
		return nil
	}
	out := new(NotifierSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotifierStatus) DeepCopyInto(out *NotifierStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotifierStatus.
func (in *NotifierStatus) DeepCopy() *NotifierStatus {
	if in == nil {
		return nil
	}
	out := new(NotifierStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackNotifier) DeepCopyInto(out *SlackNotifier) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackNotifier.
func (in *SlackNotifier) DeepCopy() *SlackNotifier {
	if in == nil {
		return nil
	}
	out := new(SlackNotifier)
	in.DeepCopyInto(out)
	return out
}
