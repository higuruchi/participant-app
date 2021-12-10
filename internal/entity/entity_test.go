package entity

import (
	"testing"
)

func TestNewParticipantEntity(t *testing.T) {
	type args struct {
		id string
		name string
	}

	type wantArgs struct {
		id string
		name string
		err error
	}

	tests := []struct {
		name string
		args args
		wantArgs wantArgs
	}{
		{
			name: "check correct operation",
			args: args{
				id: "21T325",
				name: "hoge",
			},
			wantArgs: wantArgs{
				id: "21T325",
				name: "hoge",
				err: nil,
			},
		},
		{
			name: "check correct operation",
			args: args{
				id: "20T325",
				name: "fuga",
			},
			wantArgs: wantArgs{
				id: "20T325",
				name: "fuga",
				err: nil,
			},
		},
		{
			name: "check correct operation",
			args: args{
				id: "19T325",
				name: "hogeufuga",
			},
			wantArgs: wantArgs{
				id: "19T325",
				name: "hogeufuga",
				err: nil,
			},
		},
		{
			name: "check invalid subjects",
			args: args{
				id: "19A000",
				name: "hogeufuga",
			},
			wantArgs: wantArgs{
				id: "",
				name: "",
				err: ErrInvalidId,
			},
		},
		{
			name: "check invalid id",
			args: args{
				id: "199A000",
				name: "hogeufuga",
			},
			wantArgs: wantArgs{
				id: "",
				name: "",
				err: ErrInvalidId,
			},
		},
		{
			name: "check invalid name",
			args: args{
				id: "19T000",
				name: "",
			},
			wantArgs: wantArgs{
				id: "",
				name: "",
				err: ErrInvalidName,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal, gotErr := NewParticipantEntity(tt.args.id, tt.args.name)

			if gotErr != tt.wantArgs.err {
				t.Errorf(
					"NewParticipantEntity(id string, name string) = %v, %v, want=%v",
					gotVal,
					gotErr,
					tt.wantArgs,
				)
			}
			if gotVal != nil {
				if gotVal.GetID() != tt.wantArgs.id || gotVal.GetName() != tt.wantArgs.name {
					t.Errorf(
						"NewParticipantEntity(id string, name string) = %v, %v, want=%v",
						gotVal,
						gotErr,
						tt.wantArgs,
					)
				}
			} 
		})
	}
}

func TestDistinguishGrade(t *testing.T) {
	type args struct {
		id string
		name string
	}

	type wantArgs struct {
		grade grade
		err error
	}

	tests := []struct {
		name string
		args args
		wantArgs wantArgs
	}{
		{
			name: "check correct operation",
			args: args{
				id: "21T325",
				name: "hoge",
			},
			wantArgs: wantArgs{
				grade: B1,
				err: nil,
			},
		},
		{
			name: "check correct operation",
			args: args{
				id: "20T999",
				name: "fuga",
			},
			wantArgs: wantArgs{
				grade: B2,
				err: nil,
			},
		},
		{
			name: "check correct operation",
			args: args{
				id: "19T999",
				name: "hogefuga",
			},
			wantArgs: wantArgs{
				grade: B3,
				err: nil,
			},
		},
		{
			name: "check correct operation",
			args: args{
				id: "18T999",
				name: "hogeufgahoge",
			},
			wantArgs: wantArgs{
				grade: B4,
				err: nil,
			},
		},
		{
			name: "check correct operation",
			args: args{
				id: "17T999",
				name: "hogeufgahoge",
			},
			wantArgs: wantArgs{
				grade: Error,
				err: ErrInvalidYear,
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity, _:= NewParticipantEntity(tt.args.id, tt.args.name)
			gotResult, gotErr := entity.DistinguishGrade()

			if gotResult != tt.wantArgs.grade ||
			   gotErr != tt.wantArgs.err {
				t.Errorf("DistinguishGrade() = %v %v, want = %v %v",
					gotResult,
					gotErr,
					tt.wantArgs.grade,
					tt.wantArgs.err,
				)
			}
		})
	}

	t.Log("finish test of entity.DistinguishGrade module")
}