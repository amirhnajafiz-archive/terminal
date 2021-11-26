<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Http\Response;

/**
 * Handling the user register and sign-up feature.
 * Resource controller for all users functionalities.
 *
 */
class UserController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return JsonResponse
     */
    public function index(): JsonResponse
    {
        return response()
            ->json([
                'users' => User::all()
            ]);
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param Request $request
     * @return JsonResponse
     */
    public function store(Request $request): JsonResponse
    {
        $user = User::query()
            ->create($request->all());

        return response()
            ->json([
                'name' => $user->name ?? 'None',
                'date' => $user->created_at ?? 'None'
            ]);
    }

    /**
     * Display the specified resource.
     *
     * @param $name
     * @return JsonResponse
     */
    public function show($name): JsonResponse
    {
        $user = User::query()
            ->where('name', '=', $name)
            ->first();

        return response()
            ->json([
                'status' => $user ? 'OK' : 'FAIL',
                'user' => $user->name ?? 'None'
            ]);
    }

    /**
     * Update the specified resource in storage.
     *
     * @param Request $request
     * @param  int  $id
     * @return JsonResponse
     */
    public function update(Request $request, int $id): JsonResponse
    {
        $user = User::query()
            ->findOrFail($id);

        $user->update($request->all());

        return response()
            ->json([
                'status' => 'OK',
                'name' => $user->name ?? 'None'
            ]);
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return Response
     */
    public function destroy($id)
    {
        //
    }
}
